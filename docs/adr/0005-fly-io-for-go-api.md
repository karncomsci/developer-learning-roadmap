# ใช้ Fly.io เป็นโฮสต์ของ Go API แบบเปิดเครื่องตลอด

ทุกหน้าของเว็บ render ฝั่ง server โดยเรียก Go API ถ้า API มี cold start Visitor ทุกคนจะต้องรอไปด้วย จึงเลือก Fly.io ที่ region Singapore (อยู่ใกล้ Supabase) และเปิดเครื่องตลอด (`auto_stop_machines = "off"`) ค่าใช้จ่ายโดยประมาณคือ $0 เพราะ Fly ยกเว้นบิลของบัญชีส่วนตัวที่ต่ำกว่า $5/เดือน (ตรวจเมื่อ 2026-11-18) และ deploy จาก GitHub Actions ได้ด้วย token ตัวเดียว

## Considered Options

- **Google Cloud Run**: ใช้ free tier แบบปิดเครื่องเองเมื่อไม่มีคนใช้ได้ฟรี แต่มี cold start ประมาณ 1–2 วินาที และตั้งค่า IAM ยากกว่า ถ้า Fly ยกเลิกการยกเว้นบิลต่ำกว่า $5 ตัวนี้คือตัวสำรองอันดับแรก
- **Render, Railway, Koyeb**: แผนฟรีหลับนาน (Render ตื่นช้าประมาณ 1 นาที), ดึง image แบบ private ไม่ได้ในแผนถูก หรือแผนฟรีไม่มี region ในเอเชีย

## Consequences

- image ใน GHCR ต้องเป็น public เพื่อให้ Fly ดึงได้โดยไม่ต้องใส่รหัส (repo เป็น public อยู่แล้ว และ image ไม่มี secret)
- production ของ Vercel deploy ผ่าน Vercel CLI ใน GitHub Actions แทน Git integration เพื่อรักษาลำดับ migration → api → web ส่วน Git integration ยังใช้สร้าง preview ของ PR
