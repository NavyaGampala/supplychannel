const express = require('express');
const router = express.Router();
const productController = require('../controllers/productController');

// Routes
router.post('/createProduct', productController.createProduct);
router.post('/supplyProduct', productController.supplyProduct);
router.post('/wholesaleProduct', productController.wholesaleProduct);
router.get('/queryProduct/:productID', productController.queryProduct);
router.post('/sellProduct', productController.sellProduct);

module.exports = router;
