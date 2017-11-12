class CreateIntroductions < ActiveRecord::Migration[5.1]
  def change
    create_table :introductions do |t|
      t.references  :user
      t.string      :department
      t.integer     :grade
      t.string      :comments     
      t.timestamps
    end
  end
end
