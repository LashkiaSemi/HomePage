# This file should contain all the record creation needed to seed the database with its default values.
# The data can then be loaded with the rails db:seed command (or created alongside the database with db:setup).
#
# Examples:
#
#   movies = Movie.create([{ name: 'Star Wars' }, { name: 'Lord of the Rings' }])
#   Character.create(name: 'Luke', movie: movies.first)

=begin
User.create(name: 'test1', role: 'owner', student_id: 'T314001', password: 'pass1', password_confirmation: 'pass1')
Introduction.create(user_id: 1, department: '工学部', grade: 2, comments: 'test user')

User.create(name: 'test2', role: 'member', student_id: 'T314002', password: 'pass2', password_confirmation: 'pass2')
Introduction.create(user_id: 2, department: '工学部', grade: 3, comments: 'test user')

User.create(name: 'test3', role: 'member', student_id: 'T314003', password: 'pass3', password_confirmation: 'pass3')
Introduction.create(user_id: 3, department: '工学部', grade: 4, comments: 'test user')

User.create(name: 'test4', role: 'member', student_id: 'T314004', password: 'pass4', password_confirmation: 'pass4')
Introduction.create(user_id: 4, department: '工学部', grade: 0, comments: 'test user')
=end
