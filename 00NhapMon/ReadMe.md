Python --> phân tích dữ liệu, AI
Java --> Web Back End, Android, OOP rất tốt
JavaScript --> chuyên cho Web, React Native, Desktop
(TypeScript, CoffeeScript)
----
C, C++, C#, Golang. Golang ra đời 2009 do Google phát triển. Mục tiêu của Golang là thay thế cho code base C/C++ ở Google.
Nếu viết code bằng C/C++ quá mất thời gian
Chuyển sang Python, PHP thì tốc độ lại kém.

Mục tiêu Golang
+ Nhanh, gọn như C
+ Chuyên để lập trình web back end và những tác vụ cần xử lý khối lượng tính toán lớn
+ Chuyên ứng dụng network và cloud


Các phần mềm nổi tiếng dùng Golang
+ Docker
+ Kubernetes
+ Istio
+ Traefik

Những công ty dùng Golang cho hệ thống lõi:
1- Google
2- Alibaba
3- Tiki, VinGroup, OpenCommerce, Viettel, FPT SmartCloud

Nếu dùng Java cho các ứng dụng Cloud sẽ gặp vấn đề gì?

1. Sẽ phải cài JDK rất nặng 250Mb, tốn 100Mb bộ nhớ
2. Đóng gói Java vào Docker image thì image size rất lớn, tốn nhiều RAM
3. Tốc độ thực thi Java tương đương Golang, nhưng về network app thì không bằng.
4. Spring Boot vs Fiber: tốc độ của Java Spring Boot chỉ bằng 6.2 / 100 so với Golang Fiber

Mô hình chúng ta sẽ học Golang và áp dụng sẽ là gì?

Mobile, Desktop, Web
Mobile, Web --> kết nối vào CSDL

Chia 2 loại kiến trúc:
1- Server Side Rendering: rất cổ điển, rất ổn định, rất dễ học

2- Client Side Rendering: mới, phù hợp khi phát triển song song với ứng dụng di động. Một vài framework nổi tiếng là : React, Vue, Angular, Svelte


Tại sao không chọn Node.js JavaScript mà dùng Golang?

1. JavaScript ngôn ngữ loose type, không kiểm tra kiểu chặt
2. Node.js Express triển khai mã nguồn lên server (không bảo mật)
3. Kích thước của Nodes module rất lớn, thậm chí còn lớn hơn cả Java JDK
4. Tốc độ không quá cao, chỉ bằng 1/5 so với Golang

Chọn sử dụng Golang vì:
1. Performance
2. Hướng đến kiến trúc microservice

------

go mod init khác gì với go mod tidy

go mod init duong_dan_github_repo để tạo ra một Go module

go mod tidy để kết nối Internet tải về các go package cần thiết để chạy.

Khái niệm go module giống hệt với Node.js module !

Sau khi chạy lệnh go mod init thì file go.mod sẽ được tạo ra. Bây giờ chúng ta sử dụng VSCode có thể biên dịch được

