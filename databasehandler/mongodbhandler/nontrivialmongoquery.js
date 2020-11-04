db.Personnel.find({
  "SecurityClearance": {
    $gt: 3
  },
  "Title": {
    "$in": ["Mechanic", "Biologist"]
  }
}).pretty();
