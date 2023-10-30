// Code generated by hertz generator. DO NOT EDIT.

package backend

import (
	backend "backend/biz/handler/backend"
	"github.com/cloudwego/hertz/pkg/app/server"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	root.GET("/applist", append(__pplistMw(), backend.AppList)...)
	root.POST("/login", append(_loginMw(), backend.Login)...)
	root.POST("/register", append(_registerMw(), backend.Register)...)
	{
		_api := root.Group("/api", _apiMw()...)
		_api.POST("/account", append(__dd_ccountMw(), backend.AddAccount)...)
		_api.GET("/account", append(__ccountMw(), backend.Account)...)
		_api.GET("/accounts", append(__ccountinpageMw(), backend.AccountInPage)...)
		_api.GET("/all_user", append(_usersMw(), backend.Users)...)
		_api.GET("/basic_behaviors", append(_basicbehaviorinpageMw(), backend.BasicBehaviorInPage)...)
		_api.POST("/components", append(_genecomponentMw(), backend.GeneComponent)...)
		_api.GET("/components", append(_componentinpageMw(), backend.ComponentInPage)...)
		_api.GET("/crowd", append(_crowdinpageMw(), backend.CrowdInPage)...)
		_api.GET("/crowds", append(_crowdsMw(), backend.Crowds)...)
		_api.GET("/data_sources", append(_datasourcesMw(), backend.DataSources)...)
		_api.GET("/elements", append(_elementinpageMw(), backend.ElementInPage)...)
		_api.POST("/gene_basic_behavior", append(_genebasicbehaviorMw(), backend.GeneBasicBehavior)...)
		_api.POST("/gene_rule", append(_generuleMw(), backend.GeneRule)...)
		_api.GET("/label", append(_labelinpageMw(), backend.LabelInPage)...)
		_api.GET("/labels", append(_labelsMw(), backend.Labels)...)
		_api.GET("/model", append(_modelinpageMw(), backend.ModelInPage)...)
		_api.GET("/rule_data", append(_ruledatainpageMw(), backend.RuleDataInPage)...)
		_api.GET("/rules", append(_rulesMw(), backend.Rules)...)
		_api.POST("/seq_mining", append(_seqminingMw(), backend.SeqMining)...)
		_api.GET("/seq_mining", append(_seqminingtaskinpageMw(), backend.SeqMiningTaskInPage)...)
		_api.GET("/tree_label", append(_treelabelsMw(), backend.TreeLabels)...)
		_api.GET("/users", append(_userinpageMw(), backend.UserInPage)...)
		_api.POST("/crowd", append(_crowdMw(), backend.AddCrowd)...)
		_crowd := _api.Group("/crowd", _crowdMw()...)
		_crowd.POST("/:id", append(_genecrowdMw(), backend.GeneCrowd)...)
		_crowd.DELETE("/:id", append(_deletecrowdMw(), backend.DeleteCrowd)...)
		_api.POST("/element", append(_elementMw(), backend.AddElement)...)
		_element := _api.Group("/element", _elementMw()...)
		_element.PUT("/:id", append(_updateelementMw(), backend.UpdateElement)...)
		_element.DELETE("/:id", append(_deleteelementMw(), backend.DeleteElement)...)
		{
			_group_profile := _api.Group("/group_profile", _group_profileMw()...)
			_group_profile.GET("/:id", append(_groupprofileMw(), backend.GroupProfile)...)
		}
		_api.POST("/label", append(_labelMw(), backend.AddLabel)...)
		_label := _api.Group("/label", _labelMw()...)
		_label.DELETE("/:id", append(_deletelabelMw(), backend.DeleteLabel)...)
		_label.POST("/:id", append(_genelabelMw(), backend.GeneLabel)...)
		_label.GET("/:id", append(_singlelabelMw(), backend.SingleLabel)...)
		_api.POST("/model", append(_modelMw(), backend.AddModel)...)
		_model := _api.Group("/model", _modelMw()...)
		_model.DELETE("/:id", append(_deletemodelMw(), backend.DeleteModel)...)
		_model.POST("/:id", append(_genemodelMw(), backend.GeneModel)...)
		{
			_profile := _api.Group("/profile", _profileMw()...)
			_profile.GET("/:id", append(_profile0Mw(), backend.Profile)...)
		}
		_api.POST("/rule", append(_ruleMw(), backend.AddRule)...)
		_rule := _api.Group("/rule", _ruleMw()...)
		_rule.PUT("/:id", append(_updateruleMw(), backend.UpdateRule)...)
		_rule.DELETE("/:id", append(_deleteruleMw(), backend.DeleteRule)...)
		{
			_seq_mining_result := _api.Group("/seq_mining_result", _seq_mining_resultMw()...)
			_seq_mining_result.GET("/:id", append(_seqminingresultdownloadMw(), backend.SeqMiningResultDownload)...)
		}
		_api.POST("/user", append(_userMw(), backend.AddUser)...)
		_user := _api.Group("/user", _userMw()...)
		_user.DELETE("/:id", append(_deleteuserMw(), backend.DeleteUser)...)
		{
			_upload := _user.Group("/upload", _uploadMw()...)
			_upload.POST("/:id", append(_userdatauploadMw(), backend.UserDataUpload)...)
		}
	}
}