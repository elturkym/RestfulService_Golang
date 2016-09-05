package daos;

import ("domains")



var comments [14]domains.Comment


func GetCommentsById(id int) domains.Comment {
	return comments[id - 1]
}

func getCommentsByPostId (postId int) [] domains.Comment {
	var postComments [] domains.Comment
	for i := range comments {
		if (comments[i].PostId == postId) {
			postComments = append(postComments, comments[i])
		}
	}
	return postComments;
}