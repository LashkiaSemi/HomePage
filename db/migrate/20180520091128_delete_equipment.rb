class DeleteEquipment < ActiveRecord::Migration[5.1]
  def change
    drop_table :equipment
  end
end
