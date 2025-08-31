mongosh mongodb://localhost:27017/sane_discourse --eval "
db.posts.deleteMany({});
db.users.deleteMany({});
db.reactions.deleteMany({});
print('Database reset complete');
"