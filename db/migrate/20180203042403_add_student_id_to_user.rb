class AddStudentIdToUser < ActiveRecord::Migration[5.1]
  def change
    add_column :users, :student_id, :string
  end
end
