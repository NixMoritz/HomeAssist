CREATE TABLE Receipts (
    ReceiptID INT AUTO_INCREMENT PRIMARY KEY,   -- Unique identifier for each receipt
    ReceiptNumber VARCHAR(50) NOT NULL,         -- Unique receipt number for reference
    DateIssued DATE NOT NULL,                   -- Date when the receipt was issued
    CustomerName VARCHAR(255),                  -- Name of the customer
    CustomerEmail VARCHAR(255),                 -- Email address of the customer
    TotalAmount DECIMAL(10, 2) NOT NULL,        -- Total amount of the receipt
    PaymentMethod VARCHAR(50),                  -- Payment method used (e.g., Cash, Credit Card)
    TaxAmount DECIMAL(10, 2),                   -- Amount of tax applied
    DiscountAmount DECIMAL(10, 2),              -- Discount amount applied, if any
    NetAmount DECIMAL(10, 2) NOT NULL,          -- Final amount after applying tax and discount
    Notes TEXT,                                 -- Additional notes or comments
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Timestamp for when the receipt was created
    UpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP -- Timestamp for last update
);

CREATE TABLE Stores (
    StoreID INT AUTO_INCREMENT PRIMARY KEY,    -- Unique identifier for each vendor
    VendorName VARCHAR(255) NOT NULL,           -- Name of the vendor REWE, LIDL
    Address TEXT,                               -- Address of the vendor aufschluesseln?
    UniqeUID VARCHAR(15)
    ContactPhone VARCHAR(20),                   -- Contact phone number for the vendor
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Timestamp for when the vendor was created
    UpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP -- Timestamp for last update
);

CREATE TABLE Items (
    ItemID INT AUTO_INCREMENT PRIMARY KEY,      -- Unique identifier for each item
    ItemName VARCHAR(255) NOT NULL,             -- Name or description of the item
    UnitPrice DECIMAL(10, 2) NOT NULL,          -- Price per unit of the item
    Units DECIMAL(10, 2),
    VendorID INT,                               -- Foreign key linking to the Vendors table
    Weight DECIMAL(10, 2),
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Timestamp for when the item was created
    UpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- Timestamp for last update
    FOREIGN KEY (VendorID) REFERENCES Vendors(VendorID) -- Establish relationship with Vendors table
);

CREATE TABLE ReceiptItems (
    ReceiptItemID INT AUTO_INCREMENT PRIMARY KEY, -- Unique identifier for each receipt-item entry
    ReceiptID INT,                                -- Foreign key linking to the Receipts table
    ItemID INT,                                   -- Foreign key linking to the Items table
    Quantity INT NOT NULL,                        -- Quantity of the item purchased
    TotalPrice DECIMAL(10, 2) NOT NULL,           -- Total price for the item (UnitPrice * Quantity)
    FOREIGN KEY (ReceiptID) REFERENCES Receipts(ReceiptID), -- Establish relationship with Receipts table
    FOREIGN KEY (ItemID) REFERENCES Items(ItemID)           -- Establish relationship with Items table
);