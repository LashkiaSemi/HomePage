
root = "/var/web/homepage"
state_path "#{root}/tmp/state"
activate_control_app

_app_path = "#{File.expand_path("../..", __FILE__)}"
_app_name = File.basename(_app_path)
_home = ENV.fetch("HOME") { "/var/web/homepage" }
pidfile "/var/web/homepage/tmp/#{_app_name}.pid"
bind "unix:///var/web/homepage/tmp/#{_app_name}.sock"
directory _app_path

threads_count = ENV.fetch("RAILS_MAX_THREADS") { 5 }
threads threads_count, threads_count

port        ENV.fetch("PORT") { 3000 }

environment ENV.fetch("RAILS_ENV") { "production" }

plugin :tmp_restart
