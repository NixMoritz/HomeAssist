CREATE TABLE Items2 (
    ItemID SERIAL PRIMARY KEY,                   -- Unique identifier for each item
    ItemName VARCHAR(255) NOT NULL,              -- Name or description of the item
    UnitPrice DECIMAL(10, 2) NOT NULL,           -- Price per unit of the item
    Units DECIMAL(10, 2),                        -- Units of the item
    VendorID INT,                                -- Foreign key linking to the Vendors table
    Weight DECIMAL(10, 2),                       -- Weight of the item
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Timestamp for when the item was created
    UpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Timestamp for last update
);