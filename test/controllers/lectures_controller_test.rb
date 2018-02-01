require 'test_helper'

class LecturesControllerTest < ActionDispatch::IntegrationTest
  test "should get index" do
    get lectures_index_url
    assert_response :success
  end

  test "should get new" do
    get lectures_new_url
    assert_response :success
  end

end
