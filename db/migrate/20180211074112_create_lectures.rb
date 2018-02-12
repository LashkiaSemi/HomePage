class CreateLectures < ActiveRecord::Migration[5.1]
  def change
    create_table :lectures do |t|
      t.references :user, foreign_key: true
      t.string :title
      t.string :file
      t.string :comments
      t.boolean :activation

      t.timestamps
    end
  end
end
