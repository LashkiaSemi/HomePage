class Research < ApplicationRecord
  validates :title, presence: true
  validates :author, presence: true
  validates :comments, presence: true
  
  mount_uploader :file, DocumentUploader
end
