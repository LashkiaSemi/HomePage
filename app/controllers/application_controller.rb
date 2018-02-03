class ApplicationController < ActionController::Base
  protect_from_forgery with: :exception

  ## ログイン、ログアウト関係の関数を全てのコントローラーで使用できるように
  include SessionsHelper
end
