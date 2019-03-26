//
//  ObjectiveCViewController.m
//  remove-bg
//
//  Created by Manuel Maly on 26.03.19.
//  Copyright © 2019 Kaleido AI GmbH. All rights reserved.
//

#import "ObjectiveCViewController.h"
#import <AFNetworking/AFNetworking.h>

@interface ObjectiveCViewController ()
    @property (strong, nonatomic) IBOutlet UIImageView *imageView;
    @end

@implementation ObjectiveCViewController

- (void)viewDidLoad {
    [super viewDidLoad];
    // Do any additional setup after loading the view.
}

- (void)viewDidAppear:(BOOL)animated {
    [super viewDidAppear:animated];

    [self loadImageWithFile];
}

- (void)loadImageWithURL {

    AFHTTPSessionManager *manager =
    [[AFHTTPSessionManager alloc] initWithSessionConfiguration:
     NSURLSessionConfiguration.defaultSessionConfiguration];

    manager.responseSerializer = [AFImageResponseSerializer serializer];
    [manager.requestSerializer setValue:@"INSERT_YOUR_API_KEY_HERE"
                     forHTTPHeaderField:@"X-Api-Key"];

    NSURLSessionDataTask *dataTask =
    [manager
     POST:@"https://api.remove.bg/v1.0/removebg"
     parameters:@{@"image_url": @"https://www.remove.bg/example.jpg"}
     progress:nil
     success:^(NSURLSessionDataTask * _Nonnull task, id  _Nullable responseObject) {
         if ([responseObject isKindOfClass:UIImage.class] == false) {
             return;
         }

         self.imageView.image = responseObject;
     }
     failure:^(NSURLSessionDataTask * _Nullable task, NSError * _Nonnull error) {
         // Handle error here
     }];
    [dataTask resume];
}

- (void)loadImageWithFile {

    NSURL *fileUrl = [NSBundle.mainBundle URLForResource:@"example" withExtension:@"jpg"];
    NSData *data = [NSData dataWithContentsOfURL:fileUrl];
    if (!data) {
        return;
    }

    AFHTTPSessionManager *manager =
    [[AFHTTPSessionManager alloc] initWithSessionConfiguration:
     NSURLSessionConfiguration.defaultSessionConfiguration];

    manager.responseSerializer = [AFImageResponseSerializer serializer];
    [manager.requestSerializer setValue:@"INSERT_YOUR_API_KEY_HERE"
                     forHTTPHeaderField:@"X-Api-Key"];

    NSURLSessionDataTask *dataTask =
    [manager
     POST:@"https://api.remove.bg/v1.0/removebg"
     parameters:nil
     constructingBodyWithBlock:^(id<AFMultipartFormData>  _Nonnull formData) {
         [formData appendPartWithFileData:data
                                     name:@"image_file"
                                 fileName:@"filename.jpg"
                                 mimeType:@"image/jpeg"];
     }
     progress:nil
     success:^(NSURLSessionDataTask * _Nonnull task, id  _Nullable responseObject) {
         if ([responseObject isKindOfClass:UIImage.class] == false) {
             return;
         }

         self.imageView.image = responseObject;
     }
     failure:^(NSURLSessionDataTask * _Nullable task, NSError * _Nonnull error) {
         // Handle error here
     }];

    [dataTask resume];
}

@end
