var mongoose = require('mongoose');
var Schema = mongoose.Schema;

var OrderSchema = new Schema({
    party: Number,
    deliver: {
        type: Number,
        default: 0
    },
    content: String,
    startDate: {type: Date, default: Date.now()},
    status:{
        type: String,
        enum: ['CANCELLED', 'CONFIRMED', 'ON HOLD', 'DONE'],
        default: 'ON HOLD'
    }
}); 

module.exports = mongoose.model('Order', OrderSchema);