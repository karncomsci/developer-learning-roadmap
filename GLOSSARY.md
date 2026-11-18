# Developer Roadmap

เว็บ roadmap ส่วนตัวสำหรับวางแผนและติดตามการเรียนสาย developer ของ Owner คนเดียว และเปิดให้คนอื่นดูได้

## ผู้คน

**Owner**:
เจ้าของ roadmap มีคนเดียว และเป็นคนเดียวที่สร้าง แก้ หรือลบเนื้อหาได้
_Avoid_: Admin, User

**Visitor**:
ทุกคนที่ไม่ใช่ Owner ดูเนื้อหาได้ทั้งหมดแต่แก้อะไรไม่ได้
_Avoid_: Guest, Anonymous

## เนื้อหา

**Role**:
สายอาชีพหรือเส้นทางการเรียนหนึ่งเส้น เช่น Backend Developer ภายในประกอบด้วย Topic หลายเรื่อง
_Avoid_: Track, Path, Career, Roadmap

**Topic**:
เรื่องที่ต้องเรียนหนึ่งเรื่อง ซึ่งเป็นของ Role เดียวเท่านั้น
_Avoid_: Node, Skill, Lesson, Item

**Prerequisite**:
ความสัมพันธ์ว่า Topic หนึ่งควรเรียนก่อนอีก Topic หนึ่งใน Role เดียวกัน ใช้เป็นคำแนะนำลำดับการเรียนเท่านั้น และไม่บังคับ Learning Status ของทั้งสองฝั่ง
_Avoid_: Dependency, Edge, Requirement

**Resource**:
ลิงก์ไปยังแหล่งเรียนภายนอก (บทความ, วิดีโอ, เอกสารทางการ, คอร์ส) ที่แนบไว้กับ Topic
_Avoid_: Link, Reference, Material

## ความคืบหน้า

**Learning Status**:
สถานะการเรียนของ Owner ใน Topic หนึ่ง มีสามค่า: Not started, Learning, Done แต่ละ Topic มีสถานะเป็นอิสระต่อกัน
_Avoid_: Progress, State

**Role Progress**:
สัดส่วนของ Topic ใน Role ที่มี Learning Status เป็น Done
_Avoid_: Completion, Percentage

**Ready Topic**:
Topic ที่ยังไม่ Done และ Prerequisite ทุกตัวของมัน Done แล้ว (Topic ที่ไม่มี Prerequisite นับว่าเป็น Ready Topic) ใช้ชี้ว่าควรเรียนอะไรต่อ
_Avoid_: Unmet Prerequisite, Next, Unlocked, Available
