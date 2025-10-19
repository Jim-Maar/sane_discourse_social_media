mongosh "mongodb://admin:dev_admin_password@localhost:27017/sane_discourse?authSource=admin" --eval "
db.posts.deleteMany({});
db.users.deleteMany({});
db.reactions.deleteMany({});
db.userpages.deleteMany({});
print('Database reset complete');
"