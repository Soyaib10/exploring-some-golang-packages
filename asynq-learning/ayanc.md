-
Task হলো — "কী করতে হবে" এর definition। এটা একটা data structure। এর মধ্যে থাকে:

Type — কাজটা কী ধরনের? যেমন "email:welcome"
Payload — কাজটা করতে কী কী data লাগবে? যেমন userID: 42

ব্যস। Task নিজে কোনো কাজ করে না। এটা শুধু একটা instruction/message।
Job শব্দটা সাধারণত বোঝায় — সেই task টা যখন actually execute হচ্ছে, তখন সেই running instance কে। অনেকে পুরো system কেই "job queue" বলে

-
payload কেন []byte হয়, মানে raw bytes হয়?
Go এর memory তে সুন্দর আছে। কিন্তু এটাকে Redis এ পাঠাতে হলে, network এ পাঠাতে হলে — একটা common format এ convert করতে হবে যেটা যেকোনো system বুঝবে।
সেই common format হলো bytes — raw binary data।
এখন bytes এ convert করার আগে তোমাকে একটা encoding বেছে নিতে হবে। সবচেয়ে popular হলো JSON।
তাই flow টা এরকম:

Go struct → JSON encode → []byte → Redis এ store
Redis থেকে আনো → []byte → JSON decode → Go struct

- 
ধরো তুমি একটা courier service এর কথা ভাবছো।
তুমি একটা parcel দিতে গেলে courier office এ। বললে "এই parcel টা Chittagong পাঠাও।" Courier office সেটা নিয়ে রাখল। তোমার কাজ শেষ — তুমি চলে গেলে।
এরপর courier office থেকে delivery man বের হলো, parcel তুলল, Chittagong নিয়ে গেল, deliver করল।
এখানে:

তুমি = Client
Courier office (যেখানে parcel রাখা হয়) = Redis Queue
Delivery man = Server/Worker

এখন এটাকে asynq এ map করি।
Client এর কাজ একটাই — task তৈরি করে queue তে ঢোকানো। ব্যস। সে জানে না task কখন process হবে, কে process করবে। তার দায়িত্ব শেষ।
goclient := asynq.NewClient(redisOpt)
client.Enqueue(task)  // queue তে দিয়ে দিলাম, চলে গেলাম
Server এর কাজ — queue থেকে task তুলে এনে process করা। সে সবসময় background এ চলতে থাকে, queue watch করতে থাকে। নতুন task আসলে সাথে সাথে তুলে নেয়।
gosrv := asynq.NewServer(redisOpt, config)
srv.Run(handler)  // চলতেই থাকে, থামে না

Client আর Server সম্পূর্ণ আলাদা program হতে পারে। এমনকি আলাদা machine এও থাকতে পারে। এরা একে অপরকে চেনে না — শুধু Redis কে চেনে।
Client Redis এ লেখে। Server Redis থেকে পড়ে। এটাই সংযোগ।
এই কারণে system টা scalable — তুমি চাইলে ১টা Client আর ১০টা Server চালাতে পারো। সব Server একই Redis থেকে task নেবে। কেউ কারো সাথে কথা বলবে না, race condition হবে না কারণ Redis atomic।

-
Client task দিয়ে চলে যায়। তাহলে যদি Client task দেওয়ার পরে জানতে চায় — "আমার task টা process হয়েছে কিনা?" — এটা কি সম্ভব? 
asynq হলো "fire and forget" system।
মানে Client task দিল এবং ভুলে গেল। এটাই intended behavior। কারণ পুরো point ই হলো — main flow কে block করবে না।
তুমি যদি Client কে wait করাও "task complete হয়েছে কিনা" জানার জন্য — তাহলে আর async রইল না, সেটা হয়ে গেল sync।

কিন্তু তারপরেও real world এ জানার দরকার হয়। সেটা কীভাবে হয়?
উপায় ১ — Task ID দিয়ে track করো।
Client যখন Enqueue করে, সে একটা info object পায়:
goinfo, err := client.Enqueue(task)
log.Println(info.ID)     // task এর unique ID
log.Println(info.Queue)  // কোন queue তে গেছে
log.Println(info.State)  // এখন কী state এ আছে
এই ID টা তুমি database এ save করে রাখতে পারো। পরে asynq এর Inspector দিয়ে জিজ্ঞেস করতে পারো "এই ID এর task এর কী অবস্থা?"
উপায় ২ — Callback pattern।
Task process হওয়ার পরে Handler নিজেই database এ একটা record update করে দেয়। "আমার কাজ শেষ।" Client পরে সেই database দেখে জানতে পারে।
উপায় ৩ — asynq এর Result storage।
asynq এ task result store করার feature আছে। Handler result লিখে রাখে, পরে পড়া যায়।


- 
some terms of async
Pending মানে — task queue তে আছে, কিন্তু কোনো worker এখনো তুলে নেয়নি। অপেক্ষা করছে।
Active মানে — একটা worker task টা তুলে নিয়েছে এবং এখন process করছে।
Retry → process করতে গিয়ে error হয়েছে, কিছুক্ষণ পরে আবার চেষ্টা করবে।
Archived → অনেকবার retry হয়েছে, সব fail। এখন dead। manually intervention ছাড়া আর চলবে না।
Scheduled → ভবিষ্যতে process করার জন্য দেওয়া হয়েছে। এখনো সময় হয়নি।
Completed → সফলভাবে শেষ হয়েছে।

flow: 
Scheduled (সময় হলে)
      ↓
   Pending
      ↓
   Active
    ↙    ↘
Success   Error
   ↓         ↓
Completed   Retry (আবার Pending হবে)
                ↓ (retry শেষ হলে)
             Archived
