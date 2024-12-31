/*
 * Svix API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.1.1
 * 
 * Generated by: https://openapi-generator.tech
 */

#[allow(unused_imports)]
use crate::models;
#[allow(unused_imports)]
use serde::{Deserialize, Serialize};
        
                #[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
                pub struct CompletionChoice {
                        #[serde(rename = "finish_reason")]
                        pub finish_reason: String,
                        #[serde(rename = "index")]
                        pub index: i64,
                        #[serde(rename = "message")]
                        pub message: Box<models::CompletionMessage>,
                    }

                    impl CompletionChoice {
                    pub fn new(finish_reason: String, index: i64, message: models::CompletionMessage) -> CompletionChoice {
                CompletionChoice {
                    finish_reason,
                    index,
                    message: Box::new(message),
                    }
                    }
                    }

