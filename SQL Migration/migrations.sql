CREATE TABLE
    IF NOT EXISTS Stores (
        StoreID SERIAL PRIMARY KEY, -- Unique identifier for each Store
        UniqueUID VARCHAR(15), -- Unique identifier
        StoreBranch VARCHAR(255), -- Chain name (e.g., REWE, Kaufland)
        StoreName VARCHAR(255), -- Optional name (e.g., special names like Edeka Special)
        StoreAddress TEXT, -- Address of the Store
        StorePhone VARCHAR(20), -- Contact phone number for the Store
        CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Timestamp for when the Store was created
        UpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Timestamp for last update
    );

CREATE TABLE
    IF NOT EXISTS Items (
        ItemID SERIAL PRIMARY KEY, -- Unique identifier for each item
        ItemName VARCHAR(255) NOT NULL, -- Name or description of the item
        UnitPrice DECIMAL(10, 2) NOT NULL, -- Price per unit of the item
        Units DECIMAL(10, 2) DEFAULT (1), -- Units of the item
        StoreBranch VARCHAR(255), -- Chain name (e.g., REWE, Kaufland)
        Weight DECIMAL(10, 2), -- Weight of the item
        CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Timestamp for when the item was created
        UpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Timestamp for last update
    );

CREATE TABLE
    IF NOT EXISTS Receipts (
        ReceiptID SERIAL PRIMARY KEY, -- Unique identifier for each receipt
        StoreID INT NOT NULL, -- Foreign key linking to the Vendors table
        DateIssued DATE NOT NULL, -- Date when the receipt was issued
        TotalAmount DECIMAL(10, 2) NOT NULL, -- Total amount of the receipt
        PaymentMethod VARCHAR(50), -- Payment method used (e.g., Cash, Credit Card)
        TotalDiscountAmount DECIMAL(10, 2), -- Discount amount applied, if any
        NetAmount DECIMAL(10, 2) NOT NULL, -- Final amount after applying tax and discount
        Notes TEXT, -- Additional notes or comments
        CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Timestamp for when the receipt was created
        UpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Timestamp for last update
        FOREIGN KEY (StoreID) REFERENCES Stores (StoreID) -- Establish relationship with Stores table
    );

CREATE TABLE
    IF NOT EXISTS ReceiptItems (
        ReceiptItemID SERIAL PRIMARY KEY, -- Unique identifier for each receipt-item entry
        ReceiptID INT NOT NULL, -- Foreign key linking to the Receipts table
        ItemID INT NOT NULL, -- Foreign key linking to the Items table
        Quantity INT NOT NULL, -- Quantity of the item purchased
        DiscountAmount DECIMAL(10, 2),
        TotalPrice DECIMAL(10, 2) NOT NULL, -- Total price for the item (UnitPrice * Quantity)
        FOREIGN KEY (ReceiptID) REFERENCES Receipts (ReceiptID) ON DELETE CASCADE,
        FOREIGN KEY (ItemID) REFERENCES Items (ItemID) -- Establish relationship with Items table
    );