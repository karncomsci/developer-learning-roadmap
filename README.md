# Developer Roadmap

เว็บ roadmap ส่วนตัวสำหรับวางแผนและติดตามการเรียนสาย developer
รายละเอียดอยู่ใน [docs/SPEC.md](docs/SPEC.md) คำศัพท์อยู่ใน [GLOSSARY.md](GLOSSARY.md) และการตัดสินใจอยู่ใน [docs/adr](docs/adr)

## สิ่งที่ต้องมีในเครื่อง

- Go 1.21 ขึ้นไป (`go.mod` ระบุ toolchain ไว้ Go จะดาวน์โหลดเวอร์ชันที่ถูกต้องให้เอง)
- Node.js ตาม `.nvmrc`
- Docker Desktop (ใช้กับ Supabase เท่านั้น)

## ติดตั้งครั้งแรก

```sh
npm install                 # Supabase CLI (ใช้ผ่าน npx supabase)
cd web && npm install && cp .env.example .env.local
cd ../api && cp .env.example .env
```

## รันในเครื่อง

เปิดสามหน้าต่าง:

```sh
# 1. ฐานข้อมูลและ Storage (ต้องเปิด Docker Desktop ก่อน)
npx supabase start
npx supabase db reset       # ล้างแล้วรัน migration + seed ใหม่

# 2. Go API ที่ http://localhost:8080
cd api && go run ./cmd/api

# 3. Next.js ที่ http://localhost:3000
cd web && npm run dev
```

Supabase Studio อยู่ที่ http://localhost:54323 หยุดด้วย `npx supabase stop`

## สร้างโค้ดจาก openapi.yaml

แก้ `openapi/openapi.yaml` แล้วรันทั้งสองฝั่ง (CI ตรวจว่า commit โค้ดที่ generate แล้ว):

```sh
cd api && go generate ./...
cd web && npm run gen:api
```

## Migration

```sh
npx supabase migration new <name>   # สร้างไฟล์ใน supabase/migrations
npx supabase db reset               # ทดสอบในเครื่อง
```

## ทดสอบและตรวจโค้ด

```sh
cd api && go vet ./... && go test ./...
cd web && npm run lint && npm run typecheck && npm run build
```

## Branch และการ deploy

ทำงานบน branch แล้วเปิด PR เข้า `main` ซึ่งต้องผ่าน CI ก่อน merge การ merge เข้า `main` คือการ deploy ขึ้น production (เริ่มตั้งแต่ M1)
