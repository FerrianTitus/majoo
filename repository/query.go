package repository

var (
	// Users
	GetUserAllQuery        = "SELECT * FROM Users"
	GetUserByUsernameQuery = "SELECT * FROM Users WHERE user_name = ? AND password = MD5(?)"

	// Merchants
	GetMerchantByUserIdQuery = "SELECT * FROM Merchants WHERE user_id = ?"

	// Outlets
	GetOutletByUserIdQuery = "SELECT a.outlet_name,b.merchant_name,c.user_name FROM Outlets a INNER JOIN Merchants b ON a.merchant_id=b.id INNER JOIN Users c ON b.user_id=c.id WHERE c.id=?"
)