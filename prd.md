# PRD — Multi-Workspace POS Web Application

## 1. Product Overview

Aplikasi merupakan platform **Point of Sale (POS)** berbasis web yang memungkinkan satu akun pengguna memiliki atau bergabung dengan beberapa **workspace bisnis**.

Contoh workspace:

* Coffee Shop
* Laundry
* Warung
* Restaurant
* Retail Store
* Salon
* Toko Elektronik
* Bisnis lainnya

Setiap workspace memiliki data yang terisolasi, seperti:

* Produk
* Kategori
* Transaksi
* Customer
* User
* Laporan
* Pengaturan bisnis

Satu akun dapat memiliki akses ke beberapa workspace dengan role yang berbeda.

Contoh:

> Rifky memiliki workspace `Kopi Senja` sebagai Owner dan workspace `Laundry Bersih` sebagai Admin.

---

# 2. Product Goals

### Primary Goals

1. Menyediakan sistem kasir yang mudah digunakan.
2. Mendukung berbagai jenis bisnis melalui konsep workspace.
3. Memisahkan data setiap bisnis/workspace.
4. Mendukung multi-user dan role-based access control.
5. Menyediakan dashboard penjualan untuk owner/admin.
6. Menyediakan interface kasir yang cepat untuk transaksi.
7. Web menjadi platform utama dan dapat dikembangkan menjadi mobile application.
8. Arsitektur dapat dikembangkan untuk berbagai jenis bisnis tanpa mengubah core system.

### Non-Goals untuk MVP

Belum menjadi fokus utama:

* Accounting lengkap
* Payroll
* HR Management
* Marketplace
* Online delivery
* Integrasi payment gateway kompleks
* Inventory manufacturing
* Advanced CRM

Fitur tersebut dapat dikembangkan pada fase berikutnya.

---

# 3. Target User

## 3.1 Owner

Pemilik bisnis.

Kebutuhan utama:

* Melihat performa bisnis
* Melihat penjualan
* Mengelola produk
* Mengelola karyawan
* Melihat laporan
* Mengelola workspace
* Mengatur permission

## 3.2 Admin

Pengelola operasional bisnis.

Kebutuhan:

* Mengelola produk
* Mengelola transaksi
* Mengelola customer
* Melihat laporan
* Mengelola user tertentu
* Mengatur operasional

## 3.3 User / Cashier

Pegawai yang menjalankan operasional.

Kebutuhan:

* Melakukan transaksi
* Melihat produk
* Memproses pembayaran
* Melihat transaksi yang dibuat sendiri
* Membuka/menutup shift

User tidak memiliki akses terhadap konfigurasi sensitif bisnis.

---

# 4. Core Concept: Account → Workspace → Role

Struktur utama aplikasi:

```text
Account
│
├── Workspace A
│   ├── Owner
│   ├── Admin
│   └── User / Cashier
│
├── Workspace B
│   ├── Owner
│   └── User
│
└── Workspace C
    └── Owner
```

### Account

Representasi identitas pengguna aplikasi.

Contoh:

`rifky@email.com`

Account tidak langsung memiliki transaksi.

Transaksi selalu berada di dalam workspace.

### Workspace

Representasi satu bisnis.

Contoh:

* Kopi Senja
* Laundry Bersih
* Warung Bu Sari

Setiap workspace memiliki data bisnis masing-masing.

### Membership

Relasi antara Account dan Workspace.

Contoh:

| Account | Workspace      | Role  |
| ------- | -------------- | ----- |
| Rifky   | Kopi Senja     | Owner |
| Budi    | Kopi Senja     | User  |
| Andi    | Kopi Senja     | Admin |
| Rifky   | Laundry Bersih | Owner |
| Sari    | Laundry Bersih | User  |

Konsep ini penting agar satu user dapat bekerja di banyak bisnis.

---

# 5. Role & Permission

## Owner

Role tertinggi di dalam workspace.

Permission:

* View Dashboard
* View Reports
* Manage Products
* Manage Categories
* Manage Transactions
* Manage Customers
* Manage Users
* Manage Roles
* Manage Workspace
* Manage Business Settings
* View Financial Summary

Owner dapat melakukan seluruh aktivitas workspace.

---

## Admin

Permission:

* View Dashboard
* View Reports
* Manage Products
* Manage Categories
* Manage Transactions
* Manage Customers
* Manage Users
* Manage Inventory

Admin tidak dapat:

* Menghapus workspace
* Mengubah ownership
* Mengakses konfigurasi account owner
* Mengelola billing/ownership

---

## User / Cashier

Permission:

* View POS
* Create Transaction
* Process Payment
* View Own Transactions
* Manage Customer saat transaksi
* Open Shift
* Close Shift

User tidak dapat:

* Melihat seluruh laporan bisnis
* Mengelola user
* Mengelola workspace
* Mengubah produk secara bebas
* Mengakses data keuangan sensitif

---

# 6. Workspace Selection

Setelah login, user diarahkan ke workspace selector jika memiliki lebih dari satu workspace.

Contoh:

```text
Welcome back, Rifky

Choose Workspace

┌──────────────────────┐
│ ☕ Kopi Senja        │
│ Coffee Shop          │
│ Owner                │
└──────────────────────┘

┌──────────────────────┐
│ 🧺 Laundry Bersih    │
│ Laundry               │
│ Owner                │
└──────────────────────┘

+ Create Workspace
```

Jika hanya memiliki satu workspace, aplikasi dapat langsung membuka workspace tersebut.

---

# 7. Workspace Dashboard

Dashboard berbeda berdasarkan permission user.

### Owner/Admin

Menampilkan:

* Total Sales
* Total Transactions
* Total Products
* Total Customers
* Today's Sales
* Sales Trend
* Best Selling Products
* Recent Transactions
* Payment Summary

Contoh:

```text
Today's Sales
Rp 4.250.000

Transactions
128

Customers
87

Best Seller
Es Kopi Susu
42 sold
```

Dashboard harus dapat menggunakan filter:

* Today
* Yesterday
* This Week
* This Month
* Custom Range

---

# 8. POS / Cashier

POS merupakan fitur utama untuk User/Cashier.

Flow:

```text
Open Shift
     ↓
Select Product
     ↓
Add to Cart
     ↓
Review Cart
     ↓
Select Payment
     ↓
Confirm Transaction
     ↓
Print / Digital Receipt
     ↓
Transaction Completed
```

POS harus dioptimalkan untuk transaksi cepat.

Fitur:

* Search product
* Category filter
* Product grid
* Cart
* Quantity adjustment
* Discount
* Tax
* Customer selection
* Payment method
* Transaction notes
* Receipt

---

# 9. Product Management

Owner/Admin dapat mengelola produk.

Data minimal:

* Product Name
* SKU
* Category
* Selling Price
* Cost Price
* Stock
* Unit
* Status
* Image

Contoh:

```text
Es Kopi Susu
SKU: KOP-001
Category: Coffee
Price: Rp18.000
Stock: 120
Status: Active
```

---

# 10. Category

Produk dapat dikelompokkan berdasarkan kategori.

Contoh coffee shop:

```text
Coffee
Non Coffee
Food
Snack
Add On
```

Laundry:

```text
Cuci Kering
Cuci Setrika
Setrika
Express
```

Karena kategori merupakan bagian dari workspace, setiap workspace dapat memiliki struktur kategori yang berbeda.

---

# 11. Customer

Customer management bersifat opsional tergantung jenis bisnis.

Data:

* Name
* Phone
* Email
* Address
* Transaction History
* Total Spending

Untuk coffee shop, customer bisa sederhana.

Untuk laundry, customer dapat menjadi fitur penting karena transaksi biasanya membutuhkan data pelanggan.

---

# 12. Transaction

Setiap transaksi harus memiliki:

* Transaction ID
* Workspace ID
* Cashier
* Customer
* Items
* Quantity
* Subtotal
* Discount
* Tax
* Grand Total
* Payment Method
* Payment Status
* Transaction Status
* Created At

Contoh:

```text
TRX-20260813-00125

Cashier:
Budi

Items:
2x Es Kopi Susu
1x Roti Bakar

Subtotal:
Rp52.000

Discount:
Rp2.000

Total:
Rp50.000

Payment:
Cash

Status:
Paid
```

---

# 13. Payment

MVP mendukung:

* Cash
* Bank Transfer
* QRIS
* E-Wallet

Payment architecture sebaiknya dibuat extensible agar payment provider dapat ditambahkan kemudian.

---

# 14. Shift Management

Cashier dapat membuka shift.

Contoh:

```text
Shift #001

Cashier:
Budi

Opening Balance:
Rp300.000

Opened:
08:01
```

Saat shift ditutup:

```text
Opening Balance
Rp300.000

Cash Sales
Rp1.500.000

Expected Cash
Rp1.800.000

Actual Cash
Rp1.790.000

Difference
-Rp10.000
```

Owner/Admin dapat melihat histori shift.

---

# 15. Reports

Owner/Admin dapat melihat:

### Sales Report

* Total sales
* Transaction count
* Average transaction value
* Sales by date

### Product Report

* Best seller
* Lowest seller
* Product quantity sold

### Cashier Report

* Sales by cashier
* Number of transactions
* Shift performance

### Payment Report

* Cash
* Transfer
* QRIS
* Other payment methods

---

# 16. User Management

Owner/Admin dapat mengundang user.

Flow:

```text
Workspace
   ↓
Users
   ↓
Invite User
   ↓
Email / Account
   ↓
Select Role
   ↓
Send Invitation
```

Contoh:

```text
Budi
Cashier
Active

Andi
Admin
Active

Sari
Cashier
Invited
```

---

# 17. Business Type

Workspace memiliki `business type`.

Contoh:

```text
COFFEE_SHOP
LAUNDRY
RESTAURANT
RETAIL
SALON
OTHER
```

Business type tidak boleh terlalu membatasi core POS.

Contohnya:

Coffee shop:

```text
Product → Coffee
```

Laundry:

```text
Service → Cuci Kering 5kg
```

Karena itu, secara arsitektur lebih baik core entity menggunakan konsep **Item**, yang dapat berupa:

* Product
* Service

Sehingga laundry tidak dipaksa menggunakan konsep barang/stok seperti retail.

---

# 18. Laundry Extension

Laundry dapat dikembangkan sebagai module khusus.

Contoh flow:

```text
Customer
   ↓
Laundry Order
   ↓
Select Service
   ↓
Weight / Quantity
   ↓
Price Calculation
   ↓
Payment
   ↓
Laundry Status
```

Status:

```text
Received
Processing
Ready
Picked Up
Cancelled
```

Laundry module tidak perlu masuk MVP POS core.

---

# 19. Coffee Shop Extension

Coffee shop dapat dikembangkan dengan:

* Table Management
* Order Management
* Kitchen Display
* Modifier
* Topping
* Size
* Dine-in
* Takeaway

Contoh:

```text
Es Kopi Susu

Size:
Small
Medium
Large

Add-on:
Extra Shot
Oat Milk
Ice
Sugar
```

---

# 20. Navigation

Untuk Owner/Admin:

```text
Dashboard
POS
Transactions
Products
Categories
Customers
Inventory
Reports
Users
Settings
```

Untuk Cashier:

```text
POS
Transactions
Customers
Shift
```

Menu ditampilkan berdasarkan permission.

---

# 21. Multi-Tenant Data Isolation

Workspace harus menjadi boundary utama data.

Secara konsep:

```text
workspace_id
```

harus melekat pada seluruh entity yang berhubungan dengan bisnis.

Contoh:

```text
Workspace
 ├── Products
 ├── Categories
 ├── Customers
 ├── Transactions
 ├── Users
 ├── Shifts
 └── Reports
```

User dari Workspace A tidak boleh dapat mengakses data Workspace B hanya dengan memanipulasi request/API.

Authorization wajib dilakukan di backend.

---

# 22. Mobile Application

Web merupakan platform utama pada MVP.

Mobile application dikembangkan kemudian menggunakan API yang sama.

Target penggunaan mobile:

### Owner

* Dashboard
* Sales monitoring
* Reports
* Notifications

### Cashier

* POS
* Transaction
* Customer
* Shift

Arsitektur harus API-first agar mobile tidak membutuhkan backend berbeda.

---

# 23. MVP Scope

### Phase 1 — Core Platform

* Authentication
* Account
* Workspace
* Workspace selector
* Membership
* Role & Permission
* Workspace settings

### Phase 2 — POS

* Product
* Category
* Cart
* Transaction
* Payment
* Receipt
* Shift

### Phase 3 — Management

* Customer
* User management
* Dashboard
* Sales reports

### Phase 4 — Business Modules

* Coffee Shop
* Laundry
* Restaurant
* Retail

### Phase 5 — Mobile

* Mobile Owner Dashboard
* Mobile POS
* Push Notification

---

# 24. MVP Success Criteria

MVP dianggap berhasil apabila:

1. User dapat membuat workspace.
2. Owner dapat mengundang cashier.
3. Cashier dapat melakukan transaksi.
4. Transaksi tercatat di workspace yang benar.
5. Owner dapat melihat penjualan.
6. Data antar-workspace terisolasi.
7. Permission setiap role berjalan dengan benar.
8. Sistem dapat digunakan untuk minimal dua tipe bisnis tanpa mengubah core transaction system.

---

# 25. Prinsip Arsitektur Produk

Produk harus dibangun dengan prinsip:

**Core POS + Business Modules**

Bukan:

**Coffee Shop App + fitur laundry**

Struktur yang lebih sehat:

```text
                POS Platform
                     │
        ┌────────────┼────────────┐
        │            │            │
     Retail      Coffee Shop    Laundry
        │            │            │
     Module        Module       Module
```

Dengan demikian, fitur umum seperti:

* Authentication
* Workspace
* User
* Role
* Product/Item
* Customer
* Transaction
* Payment
* Report

tetap berada di core.

Sedangkan fitur khusus bisnis menjadi module.

---

# 26. Future Expansion

Setelah core stabil, platform dapat dikembangkan menjadi:

* Inventory
* Supplier
* Purchase
* Accounting
* Loyalty
* Membership
* Subscription
* Online Ordering
* QR Ordering
* Kitchen Display
* Laundry Tracking
* Multi-Branch
* Employee Management
* Mobile Application
* Notification
* Payment Gateway
* Subscription/Billing

Untuk multi-branch, struktur dapat dikembangkan menjadi:

```text
Account
   ↓
Workspace
   ↓
Branch
   ↓
Users / POS / Transactions
```

Contoh:

```text
Kopi Senja
│
├── Cabang Sidoarjo
│   ├── POS 1
│   └── POS 2
│
└── Cabang Surabaya
    ├── POS 1
    └── POS 2
```

---

# 27. Product Vision

Aplikasi bukan sekadar aplikasi kasir.

Visinya adalah menjadi:

> **Satu platform operasional bisnis yang dapat digunakan berbagai jenis usaha melalui sistem workspace dan business module.**

User cukup memiliki satu account, kemudian dapat mengelola berbagai bisnis dari satu platform.

```text
                 ACCOUNT
                    │
        ┌───────────┼───────────┐
        ↓           ↓           ↓
    COFFEE SHOP   LAUNDRY     RETAIL
        │           │           │
       POS         POS         POS
        │           │           │
    Dashboard    Dashboard   Dashboard
```

Core platform tetap sama, sementara kebutuhan khusus tiap bisnis dikembangkan sebagai module.
