class CreateEquipments < ActiveRecord::Migration[5.1]
  def change
    create_table :equipments do |t|
      t.string :name
      t.integer :num
      t.string :note
      t.references :tag

      t.timestamps
    end
  end
end
