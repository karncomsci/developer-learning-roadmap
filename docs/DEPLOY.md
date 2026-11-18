# ตั้งค่า production ครั้งแรก

ทำครั้งเดียวก่อน merge workflow `deploy.yml` เข้า `main` ลำดับการ deploy คือ migration → api → web (SPEC §10, ADR 0005)

## 1. Supabase (ฐานข้อมูล)

1. สร้าง project ที่ https://supabase.com/dashboard โดยเลือก region **Southeast Asia (Singapore)** แล้วจดรหัสผ่านฐานข้อมูลไว้
2. จด **Project ref** (ส่วน `xxxx` ใน `https://xxxx.supabase.co`)
3. สร้าง **Access token** ที่ Account → Access Tokens
4. หา connection string สำหรับ API ที่ Connect → **Session pooler** (port 5432) แล้วใส่รหัสผ่านจริงแทน `[YOUR-PASSWORD]`

## 2. Fly.io (Go API)

ติดตั้ง `flyctl` (PowerShell: `iwr https://fly.io/install.ps1 -useb | iex`) แล้วรัน:

```powershell
fly auth login
fly apps create developer-roadmap-api        # ถ้าชื่อซ้ำ ให้เปลี่ยนชื่อใน api/fly.toml ให้ตรงกัน
fly secrets set --app developer-roadmap-api DATABASE_URL="<session pooler connection string>"
fly tokens create deploy --app developer-roadmap-api   # คัดลอก token ไปใช้ในขั้นที่ 4
```

ที่อยู่ของ API จะเป็น `https://<ชื่อ app>.fly.dev`

## 3. Vercel (Next.js)

1. Import repo ที่ https://vercel.com/new แล้วตั้ง **Root Directory = `web`**
2. Settings → Environment Variables: เพิ่ม `API_BASE_URL = https://<ชื่อ app>.fly.dev` ให้ทั้ง Production และ Preview
3. สร้าง token ที่ Account Settings → Tokens
4. หา **Team/Org ID** (Settings ของ team หรือบัญชี) และ **Project ID** (Project Settings → General)

## 4. GitHub

Settings → Secrets and variables → Actions → **New repository secret**:

| Secret | ค่า |
|---|---|
| `SUPABASE_ACCESS_TOKEN` | จากขั้น 1.3 |
| `SUPABASE_PROJECT_REF` | จากขั้น 1.2 |
| `SUPABASE_DB_PASSWORD` | รหัสผ่านฐานข้อมูลจากขั้น 1.1 |
| `FLY_API_TOKEN` | จากขั้น 2 |
| `VERCEL_TOKEN` | จากขั้น 3.3 |
| `VERCEL_ORG_ID` | จากขั้น 3.4 |
| `VERCEL_PROJECT_ID` | จากขั้น 3.4 |

Settings → Environments → **New environment** ชื่อ `production` → เปิด **Required reviewers** แล้วใส่ชื่อตัวเอง

## 5. Deploy ครั้งแรก

1. merge PR ที่มี `deploy.yml` เข้า `main`
2. ที่แท็บ Actions → Deploy → กด **Review deployments** → Approve เพื่อให้ migration รัน
3. ถ้า job `api` ล้มเหลวเพราะ Fly ดึง image ไม่ได้: ไปที่หน้า Packages ของ repo → `developer-learning-roadmap-api` → Package settings → เปลี่ยน visibility เป็น **Public** แล้วกด Re-run jobs
4. เปิด `https://<ชื่อ app>.fly.dev/healthz` และเว็บของ Vercel ต้องใช้งานได้ (production ยังไม่มีข้อมูล หน้าแรกจะขึ้นว่า "ยังไม่มี Role")
