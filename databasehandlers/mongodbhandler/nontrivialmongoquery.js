db.Personal.find({
    "security clearance":{
        $gt: 3
    },
    "position":{
        "$in":["Mechanic", "Biologist"]
    }
}).prettty();