class ApplicationController < ActionController::Base
  protect_from_forgery with: :exception

  ## ログイン、ログアウト関係の関数を全てのコントローラーで使用できるように
  include SessionsHelper

  ## ユーザーの権限確認系関数ヘルパー
  include UsersHelper
end
