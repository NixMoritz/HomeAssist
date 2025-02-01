CREATE TABLE IF NOT EXISTS STORES (
    STORE_ID SERIAL PRIMARY KEY, -- Unique identifier for each Store
    UNIQUE_UID VARCHAR(15), -- Unique identifier
    STORE_BRANCH VARCHAR(255), -- Chain name (e.g., REWE, Kaufland)
    STORE_NAME VARCHAR(255), -- Optional name (e.g., special names like Edeka Special)
    STORE_ADDRESS TEXT, -- Address of the Store
    STORE_PHONE VARCHAR(20), -- Contact phone number for the Store
    STORE_TYPE VARCHAR(50), -- Added: e.g., 'Supermarket', 'Restaurant', 'Drugstore'
    STORE_CATEGORY VARCHAR(50), -- Added: e.g., 'Grocery', 'Fast Food', 'Electronics'
    UPDATED_AT TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Timestamp for last update
);

CREATE TABLE IF NOT EXISTS ITEMS (
    ITEM_ID SERIAL PRIMARY KEY, -- Unique identifier for each item
    ITEM_NAME VARCHAR(255) NOT NULL, -- Name or description of the item
    UNIT_PRICE DECIMAL(10, 2) NOT NULL, -- Price per unit of the item
    UNITS DECIMAL(10, 2) DEFAULT (1), -- Units of the item
    STORE_BRANCH VARCHAR(255), -- Chain name (e.g., REWE, Kaufland)
    WEIGHT DECIMAL(10, 2), -- Weight of the item
    CATEGORY VARCHAR(50), -- Added: e.g., 'Produce', 'Dairy', 'Meat', 'Beverages'
    SUBCATEGORY VARCHAR(50), -- Added: e.g., 'Fruits', 'Vegetables', 'Cheese'
    IS_ORGANIC BOOLEAN DEFAULT FALSE, -- Added: Track organic products
    BRAND_NAME VARCHAR(255), -- Added: Product brand
    BARCODE VARCHAR(50), -- Added: Product barcode/SKU
    UPDATED_AT TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Timestamp for last update
);

CREATE TABLE IF NOT EXISTS RECEIPTS (
    RECEIPT_ID SERIAL PRIMARY KEY,
    STORE_ID INT NOT NULL,
    DATE_ISSUED TIMESTAMP NOT NULL, -- Changed to TIMESTAMP to store both date and time
    TOTAL_AMOUNT DECIMAL(10, 2) NOT NULL,
    PAYMENT_METHOD VARCHAR(50),
    TOTAL_DISCOUNT_AMOUNT DECIMAL(10, 2),
    NET_AMOUNT DECIMAL(10, 2) NOT NULL,
    TAX_AMOUNT DECIMAL(10, 2),
    RECEIPT_TYPE VARCHAR(50),
    LOYALTY_CARD_NUMBER VARCHAR(50),
    CASHIER_NAME VARCHAR(100),
    RECEIPT_NUMBER VARCHAR(50),
    NOTES TEXT,
    IMAGE_URL TEXT,
    UPDATED_AT TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (STORE_ID) REFERENCES STORES (STORE_ID)
);

CREATE TABLE IF NOT EXISTS RECEIPT_ITEMS (
    RECEIPT_ITEM_ID SERIAL PRIMARY KEY, -- Unique identifier for each receipt-item entry
    RECEIPT_ID INT NOT NULL, -- Foreign key linking to the Receipts table
    ITEM_ID INT NOT NULL, -- Foreign key linking to the Items table
    QUANTITY INT NOT NULL, -- Quantity of the item purchased
    UNIT_PRICE DECIMAL(10, 2) NOT NULL, -- Added: Price at time of purchase
    DISCOUNT_AMOUNT DECIMAL(10, 2),
    DISCOUNT_PERCENTAGE DECIMAL(5, 2), -- Added: Percentage of discount
    TAX_RATE DECIMAL(5, 2), -- Added: Tax rate for the item
    TAX_AMOUNT DECIMAL(10, 2), -- Added: Tax amount for the item
    TOTAL_PRICE DECIMAL(10, 2) NOT NULL, -- Total price for the item (UnitPrice * Quantity)
    WEIGHT_AT_PURCHASE DECIMAL(10, 3), -- Added: Weight if applicable (in kg)
    NOTES TEXT, -- Added: Item-specific notes
    FOREIGN KEY (RECEIPT_ID) REFERENCES RECEIPTS (RECEIPT_ID) ON DELETE CASCADE,
    FOREIGN KEY (ITEM_ID) REFERENCES ITEMS (ITEM_ID) -- Establish relationship with Items table
);

-- Added: New table for tracking promotions and deals
CREATE TABLE IF NOT EXISTS PROMOTIONS (
    PROMOTION_ID SERIAL PRIMARY KEY,
    RECEIPT_ITEM_ID INT,
    PROMOTION_TYPE VARCHAR(50), -- e.g., 'Buy One Get One', 'Percentage Off', 'Bundle'
    PROMOTION_DESCRIPTION TEXT,
    DISCOUNT_AMOUNT DECIMAL(10, 2),
    VALID_FROM DATE,
    VALID_TO DATE,
    FOREIGN KEY (RECEIPT_ITEM_ID) REFERENCES RECEIPT_ITEMS (RECEIPT_ITEM_ID)
);