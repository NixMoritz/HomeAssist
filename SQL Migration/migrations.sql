CREATE TABLE
    IF NOT EXISTS Stores (
        Store_ID SERIAL PRIMARY KEY, -- Unique identifier for each Store
        Unique_UID VARCHAR(15), -- Unique identifier
        Store_Branch VARCHAR(255), -- Chain name (e.g., REWE, Kaufland)
        Store_Name VARCHAR(255), -- Optional name (e.g., special names like Edeka Special)
        Store_Address TEXT, -- Address of the Store
        Store_Phone VARCHAR(20), -- Contact phone number for the Store
        Updated_At TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Timestamp for last update
    );

CREATE TABLE
    IF NOT EXISTS Items (
        Item_ID SERIAL PRIMARY KEY, -- Unique identifier for each item
        Item_Name VARCHAR(255) NOT NULL, -- Name or description of the item
        Unit_Price DECIMAL(10, 2) NOT NULL, -- Price per unit of the item
        Units DECIMAL(10, 2) DEFAULT (1), -- Units of the item
        Store_Branch VARCHAR(255), -- Chain name (e.g., REWE, Kaufland)
        Weight DECIMAL(10, 2), -- Weight of the item
        Updated_At TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Timestamp for last update
    );

CREATE TABLE
    IF NOT EXISTS Receipts (
        Receipt_ID SERIAL PRIMARY KEY, -- Unique identifier for each receipt
        Store_ID INT NOT NULL, -- Foreign key linking to the Vendors table
        Date_Issued DATE NOT NULL, -- Date when the receipt was issued
        Total_Amount DECIMAL(10, 2) NOT NULL, -- Total amount of the receipt
        Payment_Method VARCHAR(50), -- Payment method used (e.g., Cash, Credit Card)
        Total_Discount_Amount DECIMAL(10, 2), -- Discount amount applied, if any
        Net_Amount DECIMAL(10, 2) NOT NULL, -- Final amount after applying tax and discount
        Notes TEXT, -- Additional notes or comments
        Updated_At TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Timestamp for last update
        FOREIGN KEY (Store_ID) REFERENCES Stores (Store_ID) -- Establish relationship with Stores table
    );

CREATE TABLE
    IF NOT EXISTS ReceiptItems (
        Receipt_Item_ID SERIAL PRIMARY KEY, -- Unique identifier for each receipt-item entry
        Receipt_ID INT NOT NULL, -- Foreign key linking to the Receipts table
        Item_ID INT NOT NULL, -- Foreign key linking to the Items table
        Quantity INT NOT NULL, -- Quantity of the item purchased
        Discount_Amount DECIMAL(10, 2),
        Total_Price DECIMAL(10, 2) NOT NULL, -- Total price for the item (UnitPrice * Quantity)
        FOREIGN KEY (Receipt_ID) REFERENCES Receipts (Receipt_ID) ON DELETE CASCADE,
        FOREIGN KEY (Item_ID) REFERENCES Items (Item_ID) -- Establish relationship with Items table
    );