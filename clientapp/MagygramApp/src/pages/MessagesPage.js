import React from "react";
import PostContextProvider from "../contexts/PostContext";
import StoryContextProvider from "../contexts/StoryContext";
import CreateStoryModal from "../components/modals/CreateStoryModal";
import AddPostToFavouritesModal from "../components/modals/AddPostToFavouritesModal";
import Header from "../components/Header";
import UserContextProvider from "../contexts/UserContext";
import NotificationContextProvider from "../contexts/NotificationContext";
import UsersConversations from "../components/messages/UsersConversations";
import MessageContextProvider from "../contexts/MessageContext";
import MessageUserSearchModal from "../components/modals/MessageUserSearchModal";
import StoryMessageModal from "../components/modals/StoryMessageModal";

const MessagesPage = () => {
	return (
		<React.Fragment>
			<div>
				<UserContextProvider>
					<StoryContextProvider>
						<PostContextProvider>
							<MessageContextProvider>
								<NotificationContextProvider>
									<Header />
								</NotificationContextProvider>
								<CreateStoryModal />
								<AddPostToFavouritesModal />
								<div>
									<div className="mt-4">
										<div className="container d-flex justify-content-center">
											<div className="col-10">
												<UsersConversations />
											</div>
										</div>
									</div>
								</div>
								<StoryMessageModal />
								<MessageUserSearchModal />
							</MessageContextProvider>
						</PostContextProvider>
					</StoryContextProvider>
				</UserContextProvider>
			</div>
		</React.Fragment>
	);
};

export default MessagesPage;
