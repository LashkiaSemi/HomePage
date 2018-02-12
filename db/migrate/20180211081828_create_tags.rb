class CreateTags < ActiveRecord::Migration[5.1]
  def change
    create_table :tags do |t|
      t.references  :equipment
      t.string        :name
      t.timestamps
    end
  end
end
