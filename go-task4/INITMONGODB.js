db = new Mongo().getDB("mongodb");
db.createCollection("users", { capped: false });