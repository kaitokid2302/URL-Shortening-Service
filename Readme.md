## [https://roadmap.sh/projects/url-shortening-service](https://roadmap.sh/projects/url-shortening-service)

**1. Giới thiệu**

- Đây là 1 dự án luyện tập trên `roadmap.sh` với mục tiêu xây dựng 1 dịch vụ rút gọn URL.

**2. Đã học được những gì**

- Biết cách dùng gorm thành thạo hơn(Tham khảo [Relationship in gorm](https://github.com/harranali/gorm-relationships-examples/tree/main/has-one))

**3. Triển khai chi tiết**

- Dùng redis làm cache

  - Read through
  - Write Around
  - Write back

- Dùng kết hợp write back và write around

  - Write back dùng cron job, cứ sau 1 phút sẽ update vào database: Dùng để đồng bộ Count, số lượng người lấy shortcode
  - Write around mỗi khi có request post tạo mới 1 shortcode

- Mỗi khi có request get shortcode, đồng thời cập nhật redis để đếm count(Count trong redis mới hơn count trong database)
  - Cứ mỗi 1 phút, count được update từ redis vào database
