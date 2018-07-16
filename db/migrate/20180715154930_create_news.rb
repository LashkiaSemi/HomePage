class CreateNews < ActiveRecord::Migration[5.1]
  def change
    create_table :news do |t|
      t.string :title
      t.integer :state
      t.datetime :time

      t.timestamps
    end
  end
end
