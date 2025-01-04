CREATE TABLE IF NOT EXISTS STORES (
    STORE_ID SERIAL PRIMARY KEY, -- Unique identifier for each Store
    UNIQUE_UID VARCHAR(15), -- Unique identifier
    STORE_BRANCH VARCHAR(255), -- Chain name (e.g., REWE, Kaufland)
    STORE_NAME VARCHAR(255), -- Optional name (e.g., special names like Edeka Special)
    STORE_ADDRESS TEXT, -- Address of the Store
    STORE_PHONE VARCHAR(20), -- Contact phone number for the Store
    UPDATED_AT TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Timestamp for last update
);

CREATE TABLE IF NOT EXISTS ITEMS (
    ITEM_ID SERIAL PRIMARY KEY, -- Unique identifier for each item
    ITEM_NAME VARCHAR(255) NOT NULL, -- Name or description of the item
    UNIT_PRICE DECIMAL(10, 2) NOT NULL, -- Price per unit of the item
    UNITS DECIMAL(10, 2) DEFAULT (1), -- Units of the item
    STORE_BRANCH VARCHAR(255), -- Chain name (e.g., REWE, Kaufland)
    WEIGHT DECIMAL(10, 2), -- Weight of the item
    UPDATED_AT TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Timestamp for last update
);

CREATE TABLE IF NOT EXISTS RECEIPTS (
    RECEIPT_ID SERIAL PRIMARY KEY, -- Unique identifier for each receipt
    STORE_ID INT NOT NULL, -- Foreign key linking to the Vendors table
    DATE_ISSUED DATE NOT NULL, -- Date when the receipt was issued
    TOTAL_AMOUNT DECIMAL(10, 2) NOT NULL, -- Total amount of the receipt
    PAYMENT_METHOD VARCHAR(50), -- Payment method used (e.g., Cash, Credit Card)
    TOTAL_DISCOUNT_AMOUNT DECIMAL(10, 2), -- Discount amount applied, if any
    NET_AMOUNT DECIMAL(10, 2) NOT NULL, -- Final amount after applying tax and discount
    NOTES TEXT, -- Additional notes or comments
    UPDATED_AT TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Timestamp for last update
    FOREIGN KEY (STORE_ID) REFERENCES STORES (STORE_ID) -- Establish relationship with Stores table
);

CREATE TABLE IF NOT EXISTS RECEIPT_ITEMS (
    RECEIPT_ITEM_ID SERIAL PRIMARY KEY, -- Unique identifier for each receipt-item entry
    RECEIPT_ID INT NOT NULL, -- Foreign key linking to the Receipts table
    ITEM_ID INT NOT NULL, -- Foreign key linking to the Items table
    QUANTITY INT NOT NULL, -- Quantity of the item purchased
    DISCOUNT_AMOUNT DECIMAL(10, 2),
    TOTAL_PRICE DECIMAL(10, 2) NOT NULL, -- Total price for the item (UnitPrice * Quantity)
    FOREIGN KEY (RECEIPT_ID) REFERENCES RECEIPTS (RECEIPT_ID) ON DELETE CASCADE,
    FOREIGN KEY (ITEM_ID) REFERENCES ITEMS (ITEM_ID) -- Establish relationship with Items table
);