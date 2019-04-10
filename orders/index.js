const express = require('express');
const app = express();
const port = 3000;
const host = '0.0.0.0';
const cors = require('cors');
const axios = require('axios');

const mongoose = require('mongoose');
const Order = require('./models/order');
const router = express.Router();

let mongoDB = 'mongodb://mongodb:27017/test'
mongoose.connect(mongoDB, { useNewUrlParser: true});

let db = mongoose.connection;
db.on('error', console.error.bind(console, 'mongoDB connection error'));

app.listen(port, host, () => console.log('Example app listening on '+ port))
app.use(express.json());
app.use(cors());
 
function validate(req, res){
    /*
        gets request's headers -> token -> sends it to user's service
        allows action or abort 401
    */
    token = req.headers['authorization'].split(' ')[1];
    headers = {'Authorization': 'Bearer '+token};
    url = 'http://127.0.0.1:8000/check';

    axios.get(
        url, headers
    ).then(response => {
        res.send(response);
    })

}

router.use(function(req, res, next){

    next();
})


router.route('/order')
    .get(function(req, res){
        validate(req, res);
        // Order.find(function(err, orders){
        //     if(err)
        //         res.send(err);
        //     res.json(orders);
        // });
    })
    .post(function(req, res){
        var order = new Order(req.body);
        order.save()
            .then(user => {
                res.json(order);
            })
            .catch(err => {
                res.status(400).send('error!');
            });
    })
    .delete(function(req, res){
        Order.deleteMany({}, function(err, data){
            if(err)
                res.send(err);
            else
                res.send(data);
        })
    });


router.route('/order/:order_id')
    .delete(function(req, res){
        Order.deleteOne({
            _id: req.params.order_id
        }, function(err, order){
            if(err)
                res.send(err)
            res.json({message: 'Successfully deleted!'});
        });
    })
    .get(function(req, res){
        Order.findById(req.params.order_id, function(err, order){
            if(err)
                res.send(err);
            res.json(order);
        });
    });

router.put('/order/:order_id', function(req, res){
    res.send('puy');
});

app.use('/', router);