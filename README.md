Tender Management Backend
📚 Loyihaga umumiy ta'rif
Tender Management Backend - bu tenderlarni boshqarish tizimi, unda mijozlar tenderlarni e'lon qilishi, contractorlar esa bu tenderlarga taklif yuborishi mumkin. Loyiha foydalanuvchilarga tenderlarni yaratish, takliflarni yuborish va tenderni baholash jarayonlarini boshqarish imkonini beradi. Loyihada foydalanuvchilarni autentifikatsiya qilish, ro'lni boshqarish, tender yaratish, taklif yuborish va tenderni baholash kabi imkoniyatlar mavjud.

⚙️ Foydalanilgan texnologiyalar
Dasturlash tili: Go (Golang)
Web framework: Gin
Ma'lumotlar bazasi: PostgreSQL
API: REST API
Dokumentatsiya: Swagger
Containerization: Docker
Real-time Update: WebSocket
Keshlash: Redis
Autentifikatsiya: JWT (JSON Web Token)

🚀 API

POST/register - foydalanuvchilarni royxatga oladi, foydalanuvchi o'z rolini o'zi tanlab registratsiya qiladi
POST/login - foydalanuvchilarni tizmiga kirishi uchun access_token beradi

Clint uchun API

POST/tenders - Client uchun tender yaratish
PUT/tenders/:id - Clint yaratgan tenderini id orqali yangilashi mumkin
DELETE/tenders/:id  - Client o'z yaratgan tinderini id orqali o'chirishi mumkin
GET/tenders - Client o'zi yaratgan tinderlar ro'yxatini olishi mumkin
GET/tenders/:id/my/bids - Client o'zi yaratgan tenderini uchun yuborilgan bidlarni filterlab olishi mumkin
POST/tenders/status_change/:id/bids - Client o'z tenderiga yuborilgan bidlarni statusini o'zgartirishi mumkin ya'ni, fail yoki award qilishi mumkin
POST/tenders/:id/award/:bid_id - Client o'z tenderi uchun yuborilgan, bidlardan birini award wiladi va qolgan barcha bidlar fail bo'lib ketadi. Tenderni yakunlash
GET/users/:id/tenders - userlar o'z tenderlarini tarixini ko'radi

Contractor uchun API

POST/tenders/:id/bids - Contractor tender uchun bid yaratadi
GET/tenders/all - Contractor barcha tenderlarni ko'ra oladi
GET/tender/:id/bids - Contractor tenderning barcha bidlarini ko'ra oladi
GET/users/:id/bids - userlar o'z bidlari tarixini ko'radi


📈 Bonus Xususiyatlar
Real-time yangilanishlar

WebSocket yordamida foydalanuvchilarga real-time notificationlar yuboriladi
Rate Limiting qo'shildi. Contractor daqiqasiga bitta tender uchun 5 tadana ortiq bid yarata olmaydi
Keshlash. Barcha get funksiyalari uchun caching qo'shildi. Cachingdan tekshirib ko'ramiz agar requestga mos response bo'lsa shu responsni qaytaramiz. Agar bo'lamasa databazadan olib qaytaramiz va redisga yozib ketamiz.

🗂️ Ma'lumotlar bazasi sxemasi
User (id, username, password, role, email)
Tender (id, client_id, title, description, deadline, budget, status)
Bid (id, tender_id, contractor_id, price, delivery_time, comments, status)
Notification (id, user_id, message, relation_id, type, created_at)

🛠️ Loyihani ishga tushirish
Loyihani ishga tushirish uchun quyidagi qadamlarni bajarishingiz kerak:

1. Ma'lumotlar bazasini sozlash
Ma'lumotlar bazasini Docker yordamida ishga tushirish uchun quyidagi komandani bajarishingiz kerak:

make run_db
2. Loyihani ishga tushirish
Loyihani va boshqa zarur xizmatlarni ishga tushirish uchun quyidagi komandani bajarish kerak:

make run
Bu komanda, Golang ilovasini konteynerda ishga tushiradi va barcha kerakli xizmatlarni ishga tushuradi.

