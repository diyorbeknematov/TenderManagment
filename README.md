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

bash
Copy code
make run_db
2. Loyihani ishga tushirish
Loyihani va boshqa zarur xizmatlarni ishga tushirish uchun quyidagi komandani bajarish kerak:

bash
Copy code
make run
Bu komanda, Golang ilovasini konteynerda ishga tushiradi va barcha kerakli xizmatlarni ishga tushuradi.

🏅 Baholash Mezonalari
Avtomatik testlar: 100 bal
Real-time yangilanishlar: 40 bal
Keshlash: 30 bal
Rate Limiting: 30 bal
Jami: 200 bal
📄 API hujjatlari
Loyihada barcha API'lar uchun hujjatlar Swagger yordamida taqdim etiladi. Swagger UI orqali barcha API metodlari va ularga tegishli parametrlarni ko'rishingiz mumkin.

Swagger hujjatiga kirish uchun Swagger UI ga tashrif buyurishingiz mumkin.

🔧 Kodek sifatiga qo'yiladigan talablar
To'g'ri ishlash: Loyiha barcha funktsional talablarni qondirishi kerak.
Kod sifati: Kodni toza, modulli va yaxshi hujjatlashtirilgan holda yozish.
Xavfsizlik: Foydalanuvchi autentifikatsiyasi va ro'lga asoslangan ruxsatlar to'g'ri ishlashini ta'minlash.
Ishlash samaradorligi: Katta ma'lumotlar to'plamini samarali ishlashini ta'minlash.
📦 O'rnatish va ishlatish
Loyihani lokal kompyuteringizda o'rnatish uchun quyidagi qadamlarni bajarishingiz mumkin:

Git repozitoriyasini klonlash:
bash
Copy code
git clone https://github.com/golanguzb70/golang-compition-2024.git
cd golang-compition-2024
Zarur bo'lgan barcha kutubxonalarni o'rnatish:
bash
Copy code
go mod tidy
Loyihani ishga tushirish:
bash
Copy code
make run
💬 Aloqa
Agar sizda savollar bo'lsa yoki yordam kerak bo'lsa, biz bilan bog'lanishingiz mumkin.

GitHub Issues: Issues page
Email: support@tender-management.com
