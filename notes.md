# Notes

## init project
- [ ] create dir entity/domain

- [ ] việc tách entity ra khỏi domain nhằm mục đích sử dụng lại được các entity này cho nhiều domain khác nhau.
- [ ] create 2 simple entities: Person and Item.
- [ ] Person là thực thể đại diện cho 1 người trong tất cả các miền.
- [ ] trong valueobject  chứa những gì?
- [ ] valueobject chứa những entity không thể thay đổi được.
- [ ] valueobject trong thực tế là có chứa ID và không thể thay đổi được.
- [ ] trong code thì không chứa ID và không thể thay đổi được.
- [ ] aggregates là gì?
- [ ] là sự kết hợp của các entity và valueobject khi bạn muốn kết hợp chúng 
  - [ ] mẫu Factory -> là 1 mẫu thiết kế được sử dụng  để gói gọn logic phức tạp bên trong các hàm nhằm tạo các phiên bản mà người dùng không biết chi tiết việc triển khai thực tế bên trong.
  - 