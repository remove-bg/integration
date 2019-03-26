import Foundation
import UIKit
import Alamofire

class SwiftViewController: UIViewController {
    
    @IBOutlet var imageView: UIImageView!

    override func viewDidAppear(_ animated: Bool) {
        super.viewDidAppear(animated)

//        loadImageWithURL()
        loadImageWithFile()
    }

    func loadImageWithURL() {

        Alamofire
            .request(
                URL(string: "https://api.remove.bg/v1.0/removebg")!,
                method: .post,
                parameters: ["image_url": "https://www.remove.bg/example.jpg"],
                encoding: URLEncoding(),
                headers: [
                    "X-Api-Key": "INSERT_YOUR_API_KEY_HERE"
                ]
            )
            .responseData { imageResponse in
                guard let imageData = imageResponse.data,
                    let image = UIImage(data: imageData) else { return }

                self.imageView.image = image
            }
    }

    func loadImageWithFile() {

        guard let fileUrl = Bundle.main.url(forResource: "example", withExtension: "jpg"),
            let data = try? Data(contentsOf: fileUrl) else { return }

        Alamofire
            .upload(
                multipartFormData: { builder in
                    builder.append(
                        data,
                        withName: "image_file",
                        fileName: "filename.jpg",
                        mimeType: "image/jpeg"
                    )
                },
                to: URL(string: "https://api.remove.bg/v1.0/removebg")!,
                headers: [
                    "X-Api-Key": "INSERT_YOUR_API_KEY_HERE"
                ]
            ) { result in
                switch result {
                case .success(let upload, _, _):
                    upload.responseJSON { imageResponse in
                        guard let imageData = imageResponse.data,
                            let image = UIImage(data: imageData) else { return }

                        self.imageView.image = image
                    }
                case .failure:
                    return
                }
            }
    }


}
