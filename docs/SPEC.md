# Spec: Developer Roadmap v1

สถานะ: Owner ยืนยันแล้ว · วันที่: 2026-11-18

คำศัพท์ทั้งหมดในเอกสารนี้ใช้ตาม `GLOSSARY.md` (Role, Topic, Learning Status, Role Progress, Prerequisite, Ready Topic, Resource, Owner, Visitor) การตัดสินใจเชิงสถาปัตยกรรมอยู่ใน `docs/adr/0001`–`0005`

## 1. เป้าหมาย

เว็บ roadmap ส่วนตัวสำหรับวางแผนและติดตามการเรียนสาย developer ของ Owner คนเดียว เปิดให้ Visitor ดูได้ทั้งหมดแบบอ่านอย่างเดียว เป้าหมายรองคือฝึกทักษะทั้งฝั่ง frontend (Next.js) และ backend (Go) รวมถึง CI/CD (ADR 0001)

ฟีเจอร์หลัก 3 ข้อ:

1. Owner เพิ่ม แก้ ลบ Role ได้ และ Role แสดงเป็น card grid
2. Topic แต่ละเรื่องมี Learning Status 3 ค่า ซึ่งเป็นอิสระต่อกัน
3. Topic ที่ต่อเนื่องกันแสดงเป็นผังที่มีเส้น Prerequisite เชื่อม เพื่อชี้ว่าควรเรียนอะไรต่อ (Ready Topic) และแนบรูปประกอบได้

## 2. ขอบเขต

### ทำใน v1

- CRUD ของ Role, Topic, Resource, รูปประกอบ และ Prerequisite
- เปลี่ยน Learning Status และคำนวณ Role Progress
- ผังของ Role ที่ระบบจัดตำแหน่งเอง เน้น Ready Topic พร้อมแผงรายละเอียดด้านข้าง
- ล็อกอิน Owner ด้วย GitHub
- UI สองภาษา (ไทย, อังกฤษ)
- CI และ CD อัตโนมัติ

### ไม่ทำใน v1

- Editor ลากวางโหนดและลากเส้นในผัง (และการบันทึกตำแหน่งโหนด)
- การจัดกลุ่ม Topic
- Topic ที่แชร์ข้าม Role และ Prerequisite ข้าม Role
- หลายผู้ใช้ การสมัครสมาชิก
- Role แบบส่วนตัว
- ถังขยะ undo และประวัติการเปลี่ยนสถานะ
- มุมมองแบบรายการ การลากเรียงการ์ด การเรียง Resource หรือรูปเอง
- Redirect ของ slug เก่า การย่อรูปอัตโนมัติ การแปลเนื้อหา การตรวจลิงก์เสีย

## 3. สถาปัตยกรรม

```
เบราว์เซอร์ ──► Next.js (Vercel) ──► Go API (Docker) ──► Postgres + Storage (Supabase)
                 │
                 └─► GitHub OAuth
```

- เบราว์เซอร์คุยกับ Next.js เท่านั้น (ADR 0003)
- กฎของโดเมนและการตรวจสิทธิ์ Owner อยู่ใน Go ที่เดียว (ADR 0002)
- รูปประกอบอ่านตรงจาก URL สาธารณะของ Storage

### โครงสร้าง repo

```
/web               Next.js (App Router, TypeScript)
/api               Go (Fiber, sqlc, pgx) — module github.com/karncomsci/developer-learning-roadmap/api
/supabase          config, migrations, seed
/openapi           openapi.yaml (สัญญากลาง)
/docs              SPEC.md, adr/
/.github/workflows CI และ CD
GLOSSARY.md
README.md          คำสั่งสำหรับพัฒนาในเครื่อง
```

### เครื่องมือ

| ส่วน | เครื่องมือ |
|---|---|
| Frontend | Next.js, Tailwind CSS, next-intl, Auth.js (GitHub), React Flow + dagre |
| Backend | Go, Fiber, sqlc, pgx, oapi-codegen (โหมด fiber-server) |
| Type ฝั่ง frontend | openapi-typescript |
| ข้อมูล | Supabase Postgres + Storage, Supabase CLI ในเครื่อง |
| ทดสอบ | go test, Playwright |
| CI/CD | GitHub Actions, GHCR, Fly.io (Go API, ADR 0005), Vercel CLI + Vercel Git integration (preview) |

ใช้เวอร์ชัน stable ล่าสุด ณ วันเริ่ม M0 แล้วล็อกไว้ (`go.mod`, `package-lock.json`, `.nvmrc`) ฝั่ง Node ใช้ npm

## 4. โมเดลข้อมูล

Migration เป็นไฟล์ SQL ใน `/supabase/migrations` และเป็นแหล่งความจริงเดียวของ schema

```sql
create type learning_status as enum ('not_started', 'learning', 'done');
create type resource_type   as enum ('article', 'video', 'official_docs', 'course', 'other');

create table roles (
  id          uuid primary key default gen_random_uuid(),
  slug        text not null unique check (slug ~ '^[a-z0-9]+(-[a-z0-9]+)*$'),
  name        text not null,
  description text not null default '',
  icon        text,
  color       text,
  created_at  timestamptz not null default now(),
  updated_at  timestamptz not null default now()
);

create table topics (
  id                uuid primary key default gen_random_uuid(),
  role_id           uuid not null references roles(id) on delete cascade,
  title             text not null,
  description_md    text not null default '',
  status            learning_status not null default 'not_started',
  status_changed_at timestamptz not null default now(),
  created_at        timestamptz not null default now(),
  updated_at        timestamptz not null default now()
);

create table prerequisites (
  topic_id              uuid not null references topics(id) on delete cascade,
  prerequisite_topic_id uuid not null references topics(id) on delete cascade,
  primary key (topic_id, prerequisite_topic_id),
  check (topic_id <> prerequisite_topic_id)
);

create table resources (
  id         uuid primary key default gen_random_uuid(),
  topic_id   uuid not null references topics(id) on delete cascade,
  title      text not null,
  url        text not null,
  type       resource_type not null default 'other',
  created_at timestamptz not null default now()
);

create table topic_images (
  id           uuid primary key default gen_random_uuid(),
  topic_id     uuid not null references topics(id) on delete cascade,
  storage_path text not null unique,
  content_type text not null,
  size_bytes   integer not null,
  created_at   timestamptz not null default now()
);
```

## 5. กฎของโดเมน

ทุกข้อบังคับใช้ใน Go

### Role

- R1: slug ไม่ซ้ำ ใช้ได้เฉพาะ a-z, 0-9 และขีดกลาง
- R2: ลบ Role แล้ว Topic, Prerequisite, Resource และรูปทั้งหมดข้างในถูกลบตาม
- R3: หน้ารวมเรียงการ์ดตาม `created_at` จากเก่าไปใหม่
- R4: slug แก้ได้ ลิงก์ที่ใช้ slug เก่าตอบ 404

### Topic

- T1: Topic เป็นของ Role เดียว และย้าย Role ไม่ได้
- T2: Topic ใหม่มี Learning Status เป็น `not_started`
- T3: ลบ Topic แล้ว Prerequisite ที่ต่อกับ Topic นั้นทั้งสองทิศ รวมถึง Resource และรูป ถูกลบตาม
- T4: Resource และรูปภายใน Topic เรียงตาม `created_at` จากเก่าไปใหม่

### Learning Status

- S1: เปลี่ยนจากค่าใดไปค่าใดก็ได้
- S2: ทุกครั้งที่ค่าเปลี่ยน ให้บันทึก `status_changed_at` ถ้าส่งค่าเดิมมา `status_changed_at` ไม่เปลี่ยน
- S3: Learning Status ของแต่ละ Topic เป็นอิสระต่อกัน Prerequisite ไม่จำกัดการเปลี่ยนสถานะ

### Role Progress

- P1: เท่ากับ จำนวน Topic ที่เป็น `done` หารด้วยจำนวน Topic ทั้งหมดใน Role
- P2: Role ที่ไม่มี Topic แสดงเป็น 0/0 และ 0%
- P3: API ส่งทั้งจำนวน `done`, จำนวน `learning` และจำนวนทั้งหมด ส่วน % คำนวณฝั่งแสดงผล

### Prerequisite

- Q1: Topic ทั้งสองต้องอยู่ใน Role เดียวกัน
- Q2: ห้ามชี้ตัวเอง และห้ามซ้ำ
- Q3: ห้ามเกิดวง ถ้าคำขอทำให้เกิดวง ให้ปฏิเสธทั้งคำขอ และตอบรายชื่อ Topic ที่อยู่ในวง
- Q4: การตรวจ Q1–Q3 และการบันทึกอยู่ใน transaction เดียว โดยล็อกแถวของ Role ไว้ก่อน

### Ready Topic

- Y1: Topic เป็น Ready Topic เมื่อสถานะไม่ใช่ `done` และ Prerequisite ทุกตัวเป็น `done` (Topic ที่ไม่มี Prerequisite นับด้วย)
- Y2: เป็นค่าที่ใช้แสดงผล frontend คำนวณจากข้อมูลของ `GET /v1/roles/{slug}` ไม่เก็บใน DB

### รูปประกอบ

- I1: รับเฉพาะ png, jpg, webp โดยตรวจจากเนื้อไฟล์จริง
- I2: ขนาดไม่เกิน 4 MB ต่อรูป (ค่า `BodyLimit` เริ่มต้นของ Fiber คือ 4 MB พอดี ซึ่งรวม overhead ของ multipart ด้วย จึงต้องตั้งให้สูงกว่า เช่น 5 MB เพื่อให้ handler เป็นผู้ตอบ `file_too_large` ตามรูปแบบ error ของเรา)
- I3: เมื่อลบรูป Topic หรือ Role ให้ลบไฟล์ใน Storage ด้วย ลำดับคือเก็บรายการ path, ลบแถวใน DB, แล้วลบไฟล์ ถ้าลบไฟล์ไม่สำเร็จให้บันทึก log ไว้ (แถวใน DB ถือว่าลบแล้ว) ไฟล์ที่ค้างจัดการด้วยมือ

## 6. สิทธิ์และการยืนยันตัว

| การกระทำ | Visitor | Owner |
|---|---|---|
| ดู Role, Topic, Learning Status, Resource, รูป | ได้ | ได้ |
| สร้าง แก้ ลบ ทุกอย่าง และเปลี่ยน Learning Status | ไม่ได้ | ได้ |

ขั้นตอน:

1. Owner ล็อกอินด้วย GitHub ผ่าน Auth.js ใน Next.js
2. Next.js เก็บ GitHub user ID (ตัวเลข) ไว้ใน session
3. เมื่อเรียก endpoint ที่แก้ข้อมูล Next.js ฝั่ง server สร้าง JWT (HS256, อายุ 60 วินาที, `aud` = `roadmap-api`) แล้วแนบใน header `Authorization: Bearer`
4. Go ตรวจลายเซ็น อายุ `aud` และเทียบ `sub` กับ `OWNER_GITHUB_ID` ถ้าไม่มี token หรือ token ใช้ไม่ได้ตอบ 401 ถ้า token ถูกต้องแต่ไม่ใช่ Owner ตอบ 403
5. Endpoint แบบ GET ไม่ต้องมี token

Next.js ใช้ session เพื่อแสดงหรือซ่อนปุ่มแก้ไขเท่านั้น ผู้ตัดสินสิทธิ์จริงคือ Go

## 7. API

REST + JSON ภายใต้ `/v1` นิยามเต็มอยู่ใน `/openapi/openapi.yaml` ซึ่งเขียนก่อนลงมือโค้ด และระบุความยาวสูงสุดของชื่อและคำอธิบายไว้ที่นั่น

| Method | Path | สิทธิ์ | หน้าที่ |
|---|---|---|---|
| GET | `/healthz` | ทุกคน | ตรวจว่าบริการทำงาน |
| GET | `/v1/roles` | ทุกคน | รายการ Role พร้อมตัวเลข Role Progress |
| POST | `/v1/roles` | Owner | สร้าง Role |
| GET | `/v1/roles/{slug}` | ทุกคน | Role หนึ่งตัวพร้อม Topic, Learning Status, Prerequisite, Resource และรูปทั้งหมด |
| PATCH | `/v1/roles/{roleId}` | Owner | แก้ Role (รวม slug) |
| DELETE | `/v1/roles/{roleId}` | Owner | ลบ Role |
| POST | `/v1/roles/{roleId}/topics` | Owner | สร้าง Topic พร้อมรายการ Prerequisite |
| PATCH | `/v1/topics/{topicId}` | Owner | แก้ Topic และแทนที่รายการ Prerequisite ทั้งชุด |
| DELETE | `/v1/topics/{topicId}` | Owner | ลบ Topic |
| PUT | `/v1/topics/{topicId}/status` | Owner | ตั้ง Learning Status |
| POST | `/v1/topics/{topicId}/resources` | Owner | เพิ่ม Resource |
| PATCH | `/v1/resources/{resourceId}` | Owner | แก้ Resource |
| DELETE | `/v1/resources/{resourceId}` | Owner | ลบ Resource |
| POST | `/v1/topics/{topicId}/images` | Owner | อัปโหลดรูป (multipart) |
| DELETE | `/v1/images/{imageId}` | Owner | ลบรูป |

ไม่มี endpoint เพิ่มหรือลบ Prerequisite ทีละเส้น

### รูปแบบ error

```json
{ "code": "prerequisite_cycle", "message": "...", "details": { "cycle": ["<topicId>", "<topicId>"] } }
```

| code | HTTP | เกิดเมื่อ |
|---|---|---|
| `validation_failed` | 400 | ข้อมูลไม่ครบหรือผิดรูปแบบ |
| `unauthorized` | 401 | ไม่มี token หรือ token ใช้ไม่ได้ |
| `forbidden` | 403 | ไม่ใช่ Owner |
| `not_found` | 404 | ไม่พบสิ่งที่อ้างถึง |
| `slug_taken` | 409 | slug ซ้ำ |
| `prerequisite_cycle` | 409 | เกิดวง |
| `prerequisite_cross_role` | 422 | Prerequisite ข้าม Role |
| `file_too_large` | 413 | รูปเกิน 4 MB |
| `unsupported_media_type` | 415 | ไม่ใช่ png, jpg, webp |

## 8. Frontend

### เส้นทาง

| URL | เนื้อหา |
|---|---|
| `/` | พาไป `/th` หรือภาษาล่าสุดที่จำไว้ใน cookie |
| `/{locale}` | card grid ของ Role |
| `/{locale}/roles/{slug}` | ผังของ Role |
| `/{locale}/roles/{slug}?topic={topicId}` | ผังพร้อมเปิดแผงของ Topic นั้น |

`locale` คือ `th` หรือ `en` มีปุ่มสลับภาษาที่คงหน้าเดิมไว้

### หน้า card grid

- การ์ดแสดงชื่อ คำอธิบาย ไอคอนหรือสี แถบ Role Progress และจำนวน เช่น 8/24
- Owner เห็นปุ่มสร้าง แก้ และลบ
- ฟอร์ม Role เติม slug จากชื่อให้อัตโนมัติเมื่อชื่อเป็นอักษรอังกฤษ ถ้าเติมไม่ได้ (เช่น ชื่อภาษาไทย) ช่องว่างไว้ให้ Owner พิมพ์ และ Owner แก้ slug ได้เสมอ
- กล่องยืนยันการลบ Role บอกจำนวน Topic ที่จะหาย

### หน้าผังของ Role

- จัดตำแหน่งอัตโนมัติจากบนลงล่างด้วย dagre
- Topic ที่ไม่มี Prerequisite เชื่อมกับใครเลย จัดเป็นตารางหลายแถวไว้ด้านบนของผัง
- โหนดแสดงชื่อและสีตาม Learning Status: เทา (ยังไม่ได้ศึกษา), เหลือง (อยู่ระหว่างศึกษา), เขียว (ศึกษาสำเร็จแล้ว)
- Ready Topic มีกรอบเน้น
- ซูมและเลื่อนได้
- คลิกโหนดเปิดแผงด้านข้าง: คำอธิบาย (Markdown), Resource, รูป, Learning Status และรายการ Prerequisite
- Owner เห็นตัวเปลี่ยน Learning Status ปุ่มแก้ และปุ่มลบในแผง
- ฟอร์ม Topic มีช่อง "ต้องเรียนต่อจาก" เลือก Topic อื่นใน Role เดียวกันได้หลายตัว
- เมื่อ API ตอบ `prerequisite_cycle` ฟอร์มแสดงชื่อ Topic ที่อยู่ในวง
- กล่องยืนยันการลบ Topic บอกจำนวนเส้น Prerequisite ที่จะหาย

### การเรียกข้อมูล

- หน้าทั้งหมด render ฝั่ง server โดยเรียก Go API
- การแก้ข้อมูลทำผ่าน Server Actions ซึ่งเรียก Go API แล้วสั่งให้หน้าที่เกี่ยวข้องโหลดข้อมูลใหม่
- ตั้ง `serverActions.bodySizeLimit` ให้รองรับรูป 4 MB (ค่าเริ่มต้นคือ 1 MB และเพดานของ Vercel คือ 4.5 MB)
- Markdown ต้องผ่านการกรอง HTML ก่อนแสดง

### ข้อความ

- ข้อความ UI ทั้งหมดอยู่ในไฟล์ `messages/th.json` และ `messages/en.json`
- ป้าย Learning Status: ยังไม่ได้ศึกษา / อยู่ระหว่างศึกษา / ศึกษาสำเร็จแล้ว และ Not started / Learning / Done
- เนื้อหาของ Role, Topic และ Resource แสดงตามที่ Owner พิมพ์

## 9. ค่าตั้ง

| ตัวแปร | ใช้ที่ | ความหมาย |
|---|---|---|
| `API_BASE_URL` | web | ที่อยู่ของ Go API |
| `API_JWT_SECRET` | web, api | secret ร่วมสำหรับเซ็นและตรวจ JWT |
| `AUTH_SECRET`, `AUTH_GITHUB_ID`, `AUTH_GITHUB_SECRET` | web | Auth.js และ GitHub OAuth |
| `OWNER_GITHUB_ID` | web, api | GitHub user ID ของ Owner |
| `DATABASE_URL` | api | Postgres |
| `SUPABASE_URL`, `SUPABASE_SERVICE_KEY`, `STORAGE_BUCKET` | api | Storage (bucket แบบสาธารณะ) |

ใช้ GitHub OAuth App แยกระหว่างเครื่องพัฒนาและ production ค่า secret ของ production เก็บใน Vercel, โฮสต์ของ API และ GitHub Actions secrets

## 10. การพัฒนาและ CI/CD

### เครื่องพัฒนา (Windows)

- Docker Desktop ใช้รัน Supabase ผ่าน `supabase start` เท่านั้น
- Go รันด้วย `go run` และ Next.js รันด้วย `npm run dev` ตรงในเครื่อง
- ไม่มี task runner คำสั่งทั้งหมดเขียนไว้ใน `README.md`

### Branch

- ทำงานบน branch แล้วเปิด PR เข้า `main`
- `main` ตั้ง protection ให้ต้องผ่าน CI ก่อน merge ไม่บังคับมีผู้รีวิว
- merge เข้า `main` คือ deploy ขึ้น production

### CI (ตั้งแต่ M0, ทุก push และ PR)

- api: `go build`, `go vet`, `go test` (unit)
- web: lint และ `npm run build`
- ตรวจว่าโค้ดที่ generate จาก `openapi.yaml` และ sqlc ตรงกับที่ commit ไว้

### CD (ตั้งแต่ M1, เมื่อ merge เข้า `main`)

ลำดับ: migration → api → web

1. migration: GitHub Actions รัน `supabase db push` ไปยัง Supabase production ผ่าน GitHub Environment ที่ต้องให้ Owner กดอนุมัติทุกครั้ง
2. api: GitHub Actions build Docker image แล้ว push ขึ้น GHCR (public) จากนั้น `flyctl deploy --image` ไปยัง Fly.io region Singapore
3. web: GitHub Actions deploy production ด้วย Vercel CLI (ปิดการ deploy `main` อัตโนมัติของ Git integration ไว้เพื่อรักษาลำดับ) ส่วน PR ได้ preview deployment จาก Git integration

ขั้นตอนตั้งค่าครั้งแรกและรายการ secret อยู่ใน `docs/DEPLOY.md`

## 11. การทดสอบ

กฎ Q1–Q3 และ S2 เขียนแบบ test-first

### Go: unit

- ตรวจวง: เส้นตรง, วงสองโหนด, วงยาว, กราฟที่มีหลายทางแต่ไม่วน
- ปฏิเสธ Prerequisite ที่ชี้ตัวเอง ซ้ำ และข้าม Role
- Role Progress: ไม่มี Topic, ทุกตัว done, ผสม
- `status_changed_at` เปลี่ยนเฉพาะเมื่อค่าสถานะเปลี่ยน

### Go: integration กับ Postgres ในเครื่อง

- ทุก endpoint ที่แก้ข้อมูลตอบ 401 เมื่อไม่มี token และ 403 เมื่อ `sub` ไม่ใช่ Owner
- ลบ Role และ Topic แล้วข้อมูลลูกหายครบ
- อัปโหลดไฟล์ผิดชนิดและเกินขนาดถูกปฏิเสธ

### Frontend

- การคำนวณ Ready Topic: ไม่มี Prerequisite, Prerequisite ครบ done, ขาดบางตัว, ตัวเองเป็น done

### Playwright

- เส้นทางหลัก: ล็อกอิน → สร้าง Role → เพิ่ม Topic สองตัวพร้อม Prerequisite → เปลี่ยนสถานะ → ตรวจ Role Progress บนการ์ด
- Visitor ไม่เห็นปุ่มแก้ไข

### Seed

`/supabase/seed.sql` สร้าง Role ตัวอย่างหนึ่งตัวพร้อม Topic ราว 10 เรื่องสำหรับเครื่องพัฒนาเท่านั้น production เริ่มจากว่างและไม่รัน seed

## 12. ลำดับงาน

| ขั้น | งาน | เสร็จเมื่อ |
|---|---|---|
| M0 | `git init` + remote, ย้าย SPEC ไป `/docs`, วางโครง repo, Supabase ในเครื่อง, migration แรก, `openapi.yaml`, สร้าง type สองฝั่ง, README, CI | `supabase start` ขึ้น, Go และ Next.js build ผ่าน, `/healthz` ตอบ, CI เขียวบน PR แรก |
| M1 | ตั้ง sqlc (บน Windows ต้องรันผ่าน Docker หรือ binary เพราะต้องใช้ cgo), เลือกโฮสต์ของ API, สร้าง Supabase cloud, Dockerfile ของ API, CD ครบสามขั้น, API และหน้า card grid แบบอ่านอย่างเดียว พร้อม seed | เปิด `/th` ในเครื่องเห็นการ์ดจาก seed และเว็บจริงเปิดได้ (ว่าง) ผ่าน CD |
| M2 | ล็อกอิน Owner, JWT, CRUD ของ Role | Owner สร้าง แก้ ลบ Role ได้ และคำขอที่ไม่มี token ถูกปฏิเสธ |
| M3 | CRUD ของ Topic, Learning Status, Role Progress | เปลี่ยนสถานะแล้ว % บนการ์ดเปลี่ยนตาม |
| M4 | Prerequisite, การตรวจวง, ผัง, Ready Topic และแผงด้านข้าง | ผังวาดถูกต้อง Ready Topic ถูกเน้น และการเพิ่มเส้นที่ทำให้วนถูกปฏิเสธพร้อมบอกวง |
| M5 | Resource และรูปประกอบ | เพิ่ม ลบ Resource และรูปได้ ลบ Topic แล้วไฟล์หายจาก Storage |
| M6 | สองภาษา | ทุกข้อความ UI เปลี่ยนตาม `/th` และ `/en` |
| M7 | Playwright, ย้าย integration test และ Playwright เข้า CI | เส้นทางหลักผ่าน Playwright บน CI และเว็บจริงใช้งานได้ครบ |

M4 เป็นขั้นที่เสี่ยงที่สุด เพราะรวมการตรวจวงกับการจัดผัง

## 13. เรื่องที่ยังเปิดอยู่

- ~~โฮสต์ของ Go API~~: เลือก Fly.io แล้ว (ADR 0005)
- เงื่อนไข free tier ของ Supabase: ต้องเช็คตอนสร้าง cloud project ใน M1
- ถ้าจะนำเนื้อหาจาก roadmap.sh มาใส่ ต้องอ่านไฟล์ LICENSE ใน repo ของเขาก่อน
