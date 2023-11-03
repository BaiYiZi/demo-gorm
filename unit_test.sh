# Services
  # Ping
    # Ping
    go test -v ./services/ping/ping_test.go ./services/ping/ping.go -run TestPing

# Services
  # Single
    # Create
      # One
      go test -v ./services/single/create_test.go ./services/single/create.go -run TestCreateOne
      # Many
      go test -v ./services/single/create_test.go ./services/single/create.go -run TestCreateMany

    # Delete
      # One
      go test -v ./services/single/delete_test.go ./services/single/delete.go -run TestDeleteOne
      # Many
      go test -v ./services/single/delete_test.go ./services/single/delete.go -run TestDeleteMany

    # Retrieve
      # One
      go test -v ./services/single/retrieve_test.go ./services/single/retrieve.go -run TestRetrieveOne
      # Many
      go test -v ./services/single/retrieve_test.go ./services/single/retrieve.go -run TestRetrieveMany

    # Update
      # One
      go test -v ./services/single/update_test.go ./services/single/update.go -run TestUpdateOne
      # Many
      go test -v ./services/single/update_test.go ./services/single/update.go -run TestUpdateMany

  # Relation
    # One To One
      # Create
      go test -v ./services/relation/one_to_one_test.go ./services/relation/one_to_one.go -run TestCreateBoyAndGirl
      # Confirm
      go test -v ./services/relation/one_to_one_test.go ./services/relation/one_to_one.go -run TestConfirmBoyAndGirlRelation
      # Cancel
      go test -v ./services/relation/one_to_one_test.go ./services/relation/one_to_one.go -run TestCancelBoyAndGirlRelation
      # Select
      go test -v ./services/relation/one_to_one_test.go ./services/relation/one_to_one.go -run TestSelectGirlByBoy

    # One To Many
      # Create
      go test -v ./services/relation/one_to_many_test.go ./services/relation/one_to_many.go -run TestCreatePeopleAndPhone
      # Confirm
      go test -v ./services/relation/one_to_many_test.go ./services/relation/one_to_many.go -run TestConfirmPeopleAndPhoneRelation
      # Cancel
      go test -v ./services/relation/one_to_many_test.go ./services/relation/one_to_many.go -run TestCancelPeopleAndPhoneRelation
      # Select
      go test -v ./services/relation/one_to_many_test.go ./services/relation/one_to_many.go -run TestSelectPhoneByPeople

    # Many To Many
      # Create
      go test -v ./services/relation/many_to_many_test.go ./services/relation/many_to_many.go -run TestCreateStudentAndCourse
      # Confirm
      go test -v ./services/relation/many_to_many_test.go ./services/relation/many_to_many.go -run TestConfirmStudentAndCourseRelation
      # Cancel
      go test -v ./services/relation/many_to_many_test.go ./services/relation/many_to_many.go -run TestCancelStudentAndCourseRelation
      # Select
      go test -v ./services/relation/many_to_many_test.go ./services/relation/many_to_many.go -run TestSelectCourseInfo