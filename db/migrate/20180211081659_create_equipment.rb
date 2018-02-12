class CreateEquipment < ActiveRecord::Migration[5.1]
  def change
    create_table :equipment do |t|
      t.string  :name
      t.integer :num
      t.string  :note
      t.timestamps
    end
  end
end
