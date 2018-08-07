module Admin
  class ApplicationController < Administrate::ApplicationController
    include SessionsHelper
    include UsersHelper
    
    before_action :is_user_logged, :is_user_admin
  end
end
