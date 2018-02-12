class Lecture < ApplicationRecord
  belongs_to :user

  validates :title, presence: true
  mount_uploader :file, DocumentUploader
end
