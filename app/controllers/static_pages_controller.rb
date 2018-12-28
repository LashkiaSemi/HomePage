class StaticPagesController < ApplicationController
  def index
    @news = News.top
  end

  def activity
    @news = News.published.page(params[:page])
  end

  def publication
  end

  def job
  end

  def link
  end
end
