# This file is auto-generated from the current state of the database. Instead
# of editing this file, please use the migrations feature of Active Record to
# incrementally modify your database, and then regenerate this schema definition.
#
# Note that this schema.rb definition is the authoritative source for your
# database schema. If you need to create the application database on another
# system, you should be using db:schema:load, not running all the migrations
# from scratch. The latter is a flawed and unsustainable approach (the more migrations
# you'll amass, the slower it'll run and the greater likelihood for issues).
#
# It's strongly recommended that you check this file into your version control system.

<<<<<<< HEAD
ActiveRecord::Schema.define(version: 20180211081828) do

  create_table "documents", force: :cascade do |t|
    t.string "name"
    t.string "path"
    t.boolean "activation"
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
    t.string "comments"
  end
=======
ActiveRecord::Schema.define(version: 20180212082511) do
>>>>>>> 8366404a2dcaf167d4cdbf2ca331e3ff1766ad27

  create_table "equipment", force: :cascade do |t|
    t.string "name"
    t.integer "num"
    t.string "note"
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
  end

  create_table "introductions", force: :cascade do |t|
    t.integer "user_id"
    t.string "department"
    t.integer "grade"
    t.string "comments"
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
    t.index ["user_id"], name: "index_introductions_on_user_id"
  end

  create_table "lectures", force: :cascade do |t|
    t.integer "user_id"
    t.string "title"
    t.string "file"
    t.string "comments"
    t.boolean "activation"
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
    t.index ["user_id"], name: "index_lectures_on_user_id"
  end

  create_table "pages", force: :cascade do |t|
    t.string "title"
    t.string "contents"
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
  end

<<<<<<< HEAD
  create_table "tags", force: :cascade do |t|
    t.integer "equipment_id"
    t.string "name"
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
    t.index ["equipment_id"], name: "index_tags_on_equipment_id"
=======
  create_table "researches", force: :cascade do |t|
    t.string "title"
    t.string "author"
    t.string "file"
    t.string "comments"
    t.boolean "activation"
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
>>>>>>> 8366404a2dcaf167d4cdbf2ca331e3ff1766ad27
  end

  create_table "users", force: :cascade do |t|
    t.string "name"
    t.string "password_digest"
    t.string "role"
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
    t.string "student_id"
  end

end
