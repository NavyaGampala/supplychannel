const { connectToNetwork } = require('../fabric/gateway');

// Create Product
exports.createProduct = async (req, res) => {
    const { productID, name, description, manufacturingDate, batchNumber } = req.body;
    try {
        const { contract, gateway } = await connectToNetwork('channel2', 'supplychain', 'Producer');
        await contract.submitTransaction('CreateProduct', productID, name, description, manufacturingDate, batchNumber);
        res.status(200).send('Product created successfully.');
        await gateway.disconnect();
    } catch (error) {
        res.status(500).send(`Failed to create product: ${error.message}`);
    }
};

// Supply Product
exports.supplyProduct = async (req, res) => {
    const { productID, supplyDate, warehouseLocation } = req.body;
    try {
        const { contract, gateway } = await connectToNetwork('channel2', 'supplychain', 'Supplier');
        await contract.submitTransaction('SupplyProduct', productID, supplyDate, warehouseLocation);
        res.status(200).send('Product supplied successfully.');
        await gateway.disconnect();
    } catch (error) {
        res.status(500).send(`Failed to supply product: ${error.message}`);
    }
};

// Wholesale Product
exports.wholesaleProduct = async (req, res) => {
    const { productID, wholesaleDate, wholesaleLocation, quantity } = req.body;
    try {
        const { contract, gateway } = await connectToNetwork('channel3', 'supplychain', 'Wholeseller');
        await contract.submitTransaction('WholesaleProduct', productID, wholesaleDate, wholesaleLocation, quantity);
        res.status(200).send('Product wholesaled successfully.');
        await gateway.disconnect();
    } catch (error) {
        res.status(500).send(`Failed to wholesale product: ${error.message}`);
    }
};

// Query Product
exports.queryProduct = async (req, res) => {
    const { productID } = req.params;
    try {
        const { contract, gateway } = await connectToNetwork('channel1', 'supplychain', 'Producer');
        const result = await contract.evaluateTransaction('QueryProduct', productID);
        res.status(200).json(JSON.parse(result.toString()));
        await gateway.disconnect();
    } catch (error) {
        res.status(500).send(`Failed to query product: ${error.message}`);
    }
};

// Sell Product
exports.sellProduct = async (req, res) => {
    const { productID, buyerInfo } = req.body;
    try {
        const { contract, gateway } = await connectToNetwork('channel1', 'supplychain', 'Wholeseller');
        await contract.submitTransaction('UpdateProductStatus', productID, 'Sold', buyerInfo);
        res.status(200).send('Product sold successfully.');
        await gateway.disconnect();
    } catch (error) {
        res.status(500).send(`Failed to sell product: ${error.message}`);
    }
};
