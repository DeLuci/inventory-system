# Inventory Tracking System

This is a personalized inventory tracking system for a small business.
The system will allow accurate and faster inventory tracking for end-users.

- Built in Go version 1.19
- Uses the [chi router](github.com/go-chi/chi)
- Uses [alex edwards SCS](github.com/alexedwards/scs) session management
- Uses [asaskevich go validator](github.com/asaskevich/govalidator)
- Uses [go migrate](github.com/golang-migrate/migrate)
- Uses [viper](github.com/spf13/viper)
- Uses [nosurf](github.com/justinas/nosurf)

---
Each product has a general QRcode, so the app will use this [QR-Code scanner](github.com/mebjas/html5-qrcode).
The app will send a JSON object to the back-end, and in each object there's a code.

Example code:
```
{
    "code": "000011000731EST00721027.5"
}
```
From this code, the string will be parsed with the product ID and size.
Each value is stored in a struct model called ScanProduct.
```
ScanProduct {
    ID : "EST007210"
    size: "27.5"
}
```
The ID and size is then used in the query shown below.
```
query := `Update sizes set $1 = $1 + 1 where id =$2`
```
---
Future Features:
- Implement the web interface
- Adding analytics depending on the stakeholder needs
- Creating a Content Management System 
