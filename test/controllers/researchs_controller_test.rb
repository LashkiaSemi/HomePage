require 'test_helper'

class ResearchsControllerTest < ActionDispatch::IntegrationTest
  test "should get index" do
    get researchs_index_url
    assert_response :success
  end

  test "should get new" do
    get researchs_new_url
    assert_response :success
  end

end
