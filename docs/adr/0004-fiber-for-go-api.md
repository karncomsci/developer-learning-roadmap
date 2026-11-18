# ใช้ Fiber เป็น web framework ของ Go API

เราเลือก Fiber แทน chi หรือ `net/http` ล้วน เพราะ Owner อยากฝึกใช้ Fiber (ต่อเนื่องจากเป้าหมายใน ADR 0001) ข้อแลกเปลี่ยนคือ Fiber สร้างบน fasthttp ไม่ใช่ `net/http` ทำให้ใช้ middleware และ `httptest` ของ ecosystem มาตรฐานตรง ๆ ไม่ได้ (integration test ต้องใช้ `app.Test()`) และ handler จะผูกกับ `fiber.Ctx` วิธีลดผลกระทบคือเก็บกฎของโดเมนไว้ใน package ที่ไม่ import Fiber เลย ให้ Fiber อยู่แค่ชั้น HTTP เท่านั้น
