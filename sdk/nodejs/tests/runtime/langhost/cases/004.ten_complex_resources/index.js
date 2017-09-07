// This tests the creation of ten resources that contains "simple" input and output propeprties.
// In particular, there aren't any fancy dataflow linked properties.

let assert = require("assert");
let MyResource = require("./resource").MyResource;

for (let i = 0; i < 10; i++) {
    let name = "testResource" + i;
    let res = new MyResource(name, i);
    res.id.mapValue(id => {
        console.log(`${name}.ID: ${id}`);
        assert.equal(id, name);
    });
    res.urn.mapValue(urn => {
        console.log(`${name}.URN: ${urn}`);
        assert.equal(urn, "test:index:MyResource::" + name);
    });
    res.outseq.mapValue(seq => {
        console.log(`${name}.Seq: ${seq}`);
        assert.equal(seq, i);
    });
    res.outprop1.mapValue(prop => {
        console.log(`${name}.OutProp1: ${prop}`);
        assert.equal(prop, "output properties ftw");
    });
    res.outprop2.mapValue(prop => {
        console.log(`${name}.OutProp2: ${prop}`);
        assert.equal(prop, 998.6);
    });
}
