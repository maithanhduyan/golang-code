import psycopg2
from faker import Faker

# Kết nối đến cơ sở dữ liệu PostgreSQL
conn = psycopg2.connect(
    host="localhost",
    port="5432",
    database="go-assets",
    user="postgres",
    password="123"
)

# Tạo con trỏ cursor để thực thi truy vấn SQL
cur = conn.cursor()

# Sử dụng thư viện Faker để tạo dữ liệu mẫu
fake = Faker()

# Tạo 100 assets với dữ liệu mẫu
for _ in range(10000):
    name = fake.name()
    code = str(fake.random_int(10000000, 99999999)).upper()  # Sửa đổi dòng này
    price = fake.pyfloat(min_value=0, max_value=100, right_digits=2)
    website = fake.url()
    description = fake.text()

    # Tạo truy vấn INSERT để chèn dữ liệu vào bảng assets
    query = f"INSERT INTO assets (name, code, price, website, description) VALUES ('{name}', '{code}', {price}, '{website}', '{description}')"
    
    # Thực thi truy vấn
    cur.execute(query)

# Lưu các thay đổi vào cơ sở dữ liệu
conn.commit()

# Đóng kết nối
cur.close()
conn.close()
