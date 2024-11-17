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

POST/tenders/


Foydalanuvchilarni ro'yxatdan o'tkazish va tizimga kirish.
Mijozlar tender e'lon qilishlari, pudratchilar esa tenderlarga taklif yuborishlari mumkin.
JWT yordamida autentifikatsiya va ro'lga asoslangan ruxsatlar.
Tender yaratish (Mijozlar uchun)

Mijozlar yangi tender yaratishlari mumkin.
Tenderlar turli holatlarda (ochiq, yopiq, taqdim etilgan) bo'lishi mumkin.
Mijozlar barcha tenderlarini ko'rish va ularning holatini boshqarishlari mumkin.
Taklif yuborish (Pudratchilar uchun)

Pudratchilar ochiq tenderlarga taklif yuborishlari mumkin.
Takliflar narx, yetkazib berish vaqti va izohlarni o'z ichiga olishi mumkin.
Takliflarni filtrlash va saralash

Mijozlar takliflarni narx va yetkazib berish vaqti bo'yicha saralashlari mumkin.
Tenderni baholash va g'olibni e'lon qilish

Mijozlar tender muddati tugaganidan so'ng barcha takliflarni baholashlari va g'olibni tanlashlari mumkin.
G'olib pudratchiga xabar yuboriladi va tender holati "yopiq" yoki "berilgan" ga o'tkaziladi.
Tender va Takliflar Tarixi

Mijozlar va pudratchilar o'zlarining yaratgan tenderlarini va yuborgan takliflarini ko'rishlari mumkin.
📈 Bonus Xususiyatlar (Ixtiyoriy)
Real-time yangilanishlar

WebSocket yordamida foydalanuvchilarga real-time yangilanishlar yuboriladi.
Rate Limiting

Taklif yuborish endpointiga rate-limiting qo'yiladi, bunda har bir pudratchiga minutiga 5 ta taklif yuborish mumkin.
Keshlash

Keshlash yordamida tenderlar va takliflarni tezroq olish imkoniyati yaratiladi.
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


📄 API hujjatlari
Loyihada barcha API'lar uchun hujjatlar Swagger yordamida taqdim etiladi. Swagger UI orqali barcha API metodlari va ularga tegishli parametrlarni ko'rishingiz mumkin.

Swagger hujjatiga kirish uchun Swagger UI ga tashrif buyurishingiz mumkin.

Git repozitoriyasini klonlash:
git clone https://github.com/golanguzb70/golang-compition-2024.git
cd golang-compition-2024
Zarur bo'lgan barcha kutubxonalarni o'rnatish:
go mod tidy
Loyihani ishga tushirish:
make run
