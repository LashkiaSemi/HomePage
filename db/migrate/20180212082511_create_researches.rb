class CreateResearches < ActiveRecord::Migration[5.1]
  def change
    create_table :researches do |t|
      t.string :title
      t.string :author
      t.string :file
      t.string :comments
      t.boolean :activation

      t.timestamps
    end
  end
end
